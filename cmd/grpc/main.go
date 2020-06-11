package main

import (
	"github.com/brgv/sv/gen/rpc"
	"github.com/brgv/sv/internal/rpc/person"
	"github.com/brgv/sv/internal/rpc/position"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	rpc.RegisterPersonServiceServer(server, &person.RpcServer{})
	rpc.RegisterPositionServiceServer(server, &position.RpcServer{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
