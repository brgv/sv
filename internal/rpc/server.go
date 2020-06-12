package rpc

import (
	"github.com/brgv/sv/gen/rpc"
	"github.com/brgv/sv/internal/rpc/service/person"
	"github.com/brgv/sv/internal/rpc/service/position"
	"google.golang.org/grpc"
	"net"
)

func NewListener(address string) (net.Listener, error) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	return listener, nil
}

func NewServer() (*grpc.Server, error) {
	server := *grpc.NewServer()
	return &server, nil
}

func SetupServer(server *grpc.Server) error {

	if serviceServer, err := person.NewServiceServer(); err == nil {
		rpc.RegisterPersonServiceServer(server, serviceServer)
	} else {
		return err
	}

	if serviceServer, err := position.NewServiceServer(); err != nil {
		rpc.RegisterPositionServiceServer(server, serviceServer)
	} else {
		return err
	}

	return nil
}
