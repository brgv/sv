package department

import (
	"context"
	"github.com/brgv/sv/gen/rpc"
)

func NewServiceServer() (rpc.DepartmentServiceServer, error) {
	return &serviceServer{}, nil
}

type serviceServer struct{}

func (s *serviceServer) Create(ctx context.Context, in *rpc.DepartmentRequest) (*rpc.DepartmentResponse, error) {
	return nil, nil
}

func (s *serviceServer) Update(ctx context.Context, in *rpc.DepartmentRequest) (*rpc.DepartmentResponse, error) {
	return nil, nil
}

func (s *serviceServer) Destroy(ctx context.Context, in *rpc.IdentityRequest) (*rpc.StatusResponse, error) {
	return nil, nil
}

func (s *serviceServer) Retrieve(ctx context.Context, in *rpc.IdentityRequest) (*rpc.DepartmentResponse, error) {
	return nil, nil
}

func (s *serviceServer) Read(ctx context.Context, in *rpc.ListRequest) (*rpc.ListDepartmentResponse, error) {
	return nil, nil
}
