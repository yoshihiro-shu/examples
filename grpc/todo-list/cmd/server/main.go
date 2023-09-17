package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/yoshihiro-shu/examples/grpc/todo-list/todo"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTodoServiceServer
}

func (s *server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Title:       in.GetTitle(),
			Description: in.GetDescription(),
		},
	}, nil
}

func (s *server) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return &pb.ListTasksResponse{}, nil
}

func (s *server) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	return &pb.GetTaskResponse{}, nil
}

func (s *server) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	return &pb.UpdateTaskResponse{}, nil
}

func (s *server) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return &pb.DeleteTaskResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
