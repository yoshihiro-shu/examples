package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/yoshihiro-shu/examples/grpc/todo-list/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "todo"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "todo client")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	r, err := c.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "test title 1",
		Description: "test descrption 1",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	t := r.GetTask()
	log.Printf("task: %s", t.Title)
	log.Printf("task: %s", t.Description)
}
