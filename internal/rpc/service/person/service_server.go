package person

import (
	"context"
	"github.com/brgv/sv/gen/rpc"
)

func NewServiceServer() (rpc.PersonServiceServer, error) {
	return &serviceServer{}, nil
}

type serviceServer struct{}

func (s *serviceServer) Create(ctx context.Context, in *rpc.PersonRequest) (*rpc.PersonResponse, error) {
	return nil, nil
}

func (s *serviceServer) Update(ctx context.Context, in *rpc.PersonRequest) (*rpc.PersonResponse, error) {
	return nil, nil
}

func (s *serviceServer) Destroy(ctx context.Context, in *rpc.IdentityRequest) (*rpc.StatusResponse, error) {
	return nil, nil
}

func (s *serviceServer) Retrieve(ctx context.Context, in *rpc.IdentityRequest) (*rpc.PersonResponse, error) {
	return nil, nil
}

func (s *serviceServer) Read(ctx context.Context, in *rpc.ListRequest) (*rpc.ListPersonResponse, error) {
	return nil, nil
}

func (s *serviceServer) Find(ctx context.Context, in *rpc.ListRequest) (*rpc.ListItemResponse, error) {
	return nil, nil
}
