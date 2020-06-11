package person

import (
	"context"
	"github.com/brgv/sv/gen/rpc"
)

type RpcServer struct{}

func (s *RpcServer) GetPersonList(ctx context.Context, in *rpc.PersonListFilter) (*rpc.PersonList, error) {
	return nil, nil
}

func (s *RpcServer) GetPerson(ctx context.Context, in *rpc.PersonId) (*rpc.Person, error) {
	return nil, nil
}

func (s *RpcServer) CreatePerson(ctx context.Context, in *rpc.PersonData) (*rpc.Person, error) {
	return nil, nil
}

func (s *RpcServer) UpdatePerson(ctx context.Context, in *rpc.Person) (*rpc.Person, error) {
	return nil, nil
}

func (s *RpcServer) DestroyPerson(ctx context.Context, in *rpc.PersonId) (*rpc.Person, error) {
	return nil, nil
}
