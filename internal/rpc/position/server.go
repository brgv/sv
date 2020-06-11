package position

import (
	"context"
	"github.com/brgv/sv/gen/rpc"
)

type RpcServer struct{}

func (s *RpcServer) GetPositionList(ctx context.Context, in *rpc.PositionListFilter) (*rpc.PositionList, error) {
	return nil, nil
}

func (s *RpcServer) GetPosition(ctx context.Context, in *rpc.PositionId) (*rpc.Position, error) {
	return nil, nil
}

func (s *RpcServer) CreatePosition(ctx context.Context, in *rpc.PositionData) (*rpc.Position, error) {
	return nil, nil
}

func (s *RpcServer) UpdatePosition(ctx context.Context, in *rpc.Position) (*rpc.Position, error) {
	return nil, nil
}

func (s *RpcServer) DestroyPosition(ctx context.Context, in *rpc.PositionId) (*rpc.Position, error) {
	return nil, nil
}
