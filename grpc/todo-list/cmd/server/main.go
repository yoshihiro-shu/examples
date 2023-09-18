package main

import (
	"context"
	"flag"
	"log"

	"github.com/yoshihiro-shu/examples/grpc/todo-list/server"
)

var (
	addr    = flag.String("addr", ":50051", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := server.Run(ctx, *network, *addr); err != nil {
		log.Fatal(err)
	}
}
