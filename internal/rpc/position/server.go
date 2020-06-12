package position

import (
	"context"
	"github.com/brgv/sv/gen/rpc"
)

func NewServiceServer() (rpc.PositionServiceServer, error) {
	return &serviceServer{}, nil
}

type serviceServer struct{}

func (s *serviceServer) Create(ctx context.Context, in *rpc.PositionRequest) (*rpc.PositionResponse, error) {
	return nil, nil
}

func (s *serviceServer) Update(ctx context.Context, in *rpc.PositionRequest) (*rpc.PositionResponse, error) {
	return nil, nil
}

func (s *serviceServer) Destroy(ctx context.Context, in *rpc.IdentityRequest) (*rpc.StatusResponse, error) {
	return nil, nil
}

func (s *serviceServer) Retrieve(ctx context.Context, in *rpc.IdentityRequest) (*rpc.PositionResponse, error) {
	return nil, nil
}

func (s *serviceServer) Read(ctx context.Context, in *rpc.ListRequest) (*rpc.ListPositionResponse, error) {
	return nil, nil
}
