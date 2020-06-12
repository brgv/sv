package main

import (
	"github.com/brgv/sv/internal/rpc"
	"google.golang.org/grpc/reflection"
	"log"
)

const (
	port = ":50051"
)

func main() {

	server, err := rpc.NewServer()

	log.Printf("Error: %#v", err)
	log.Printf("Server: %#v", server)

	if err != nil {
		panic(err)
	}

	rpc.SetupServer(server)

	reflection.Register(server)

	if listener, err := rpc.NewListener(port); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}

}
