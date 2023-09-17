package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/yoshihiro-shu/examples/grpc/todo-list/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	type args struct {
		Title       string
		Description string
	}

	type expect struct {
		Title       string
		Description string
	}

	tests := []struct {
		name   string
		input  args
		expect expect
	}{
		{
			name: "test1",
			input: args{
				Title:       "test title 1",
				Description: "test descrption 1",
			},
			expect: expect{
				Title:       "test title 1",
				Description: "test descrption 1",
			},
		},
		{
			name: "test2",
			input: args{
				Title:       "test title 2",
				Description: "test descrption 2",
			},
			expect: expect{
				Title:       "test title 2",
				Description: "test descrption 2",
			},
		},
	}

	for _, test := range tests {
		resp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{
			Title:       test.input.Title,
			Description: test.input.Description,
		})
		if err != nil {
			t.Fatal(err)
		}
		if resp.GetTask().Title != test.expect.Title {
			t.Fatalf("hello reply must be '%s'", test.expect.Title)
		}

		if resp.GetTask().Description != test.expect.Description {
			t.Fatalf("hello reply must be '%s'", test.expect.Description)
		}
	}
}
