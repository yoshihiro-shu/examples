package server

import (
	"context"

	pb "github.com/yoshihiro-shu/examples/grpc/todo-list/todo"
)

type taskServer struct {
	pb.UnimplementedTodoServiceServer
}

func newTaskServer() *taskServer {
	return new(taskServer)
}

func (s *taskServer) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Title:       in.GetTitle(),
			Description: in.GetDescription(),
		},
	}, nil
}

func (s *taskServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return &pb.ListTasksResponse{}, nil
}

func (s *taskServer) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	return &pb.GetTaskResponse{}, nil
}

func (s *taskServer) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	return &pb.UpdateTaskResponse{}, nil
}

func (s *taskServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return &pb.DeleteTaskResponse{}, nil
}
