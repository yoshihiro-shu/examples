package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var addr = []string{
	"localhost:2379",
}

func main() {
	cli, err := NewEtcdClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer cli.Close()

	err = cli.Put("hoge", "hogehoge")
	if err != nil {
		log.Fatal(err)
		return
	}
	gRes, err := cli.Get("hoge")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range gRes {
		v.Print()
	}
}

type EtcdClient struct {
	*clientv3.Client
}

func (ec *EtcdClient) Put(key, value string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := ec.Client.Put(ctx, key, value, opts...)
	cancel()
	if err != nil {
		var message string
		switch err {
		case context.Canceled:
			message = fmt.Sprintf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			message = fmt.Sprintf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			message = fmt.Sprintf("client-side error: %v\n", err)
		default:
			message = fmt.Sprintf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
		return errors.New(message)
	}
	return nil
}

func (ec *EtcdClient) Get(key string, opts ...clientv3.OpOption) ([]GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	r, err := ec.Client.Get(ctx, key, opts...)
	cancel()
	if err != nil {
		return nil, err
	}

	res := make([]GetResponse, 0)
	for _, v := range r.Kvs {
		res = append(res, GetResponse{
			Key:   v.Key,
			Value: v.Value,
		})
	}
	return res, nil
}

func NewEtcdClient(addr []string) (*EtcdClient, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &EtcdClient{c}, nil
}

type GetResponse struct {
	Key, Value []byte
}

func (gr GetResponse) Print() {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Key: %s\n", gr.Key)
	fmt.Printf("Value: %s\n", gr.Value)
}
