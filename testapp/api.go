package testapp

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/videocoin/go-bridge/testapp/types"

	"github.com/sirupsen/logrus"
)

func Register(gsrv *grpc.Server, txserver types.TransfersServiceServer) {
	types.RegisterTransfersServiceServer(gsrv, txserver)
}

func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, txserver types.TransfersServiceServer) error {
	return types.RegisterTransfersServiceHandlerServer(ctx, mux, txserver)
}

func NewServer(log *logrus.Entry, limit int64, db *DB) *Server {
	return &Server{
		log:   log,
		limit: limit,
		db:    db,
	}
}

type Server struct {
	limit int64
	log   *logrus.Entry
	db    *DB
}

func (s *Server) GetTransfers(ctx context.Context, in *types.TransfersRequest) (*types.TransfersResponse, error) {
	return nil, nil
}
