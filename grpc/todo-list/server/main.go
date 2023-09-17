package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yoshihiro-shu/examples/grpc/todo-list/todo"
)

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, newTaskServer())
	log.Printf("server listening at %v", l.Addr())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}
