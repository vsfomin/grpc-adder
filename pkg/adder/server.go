package adder

import (
	"context"

	"github.com/vsfomin/grpc-adder/api/proto"
)

// GRPC server...
type GRPCServer struct {
	proto.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) SayHello(ctx context.Context, req *proto.Message) (*proto.Message, error) {
	return &proto.Message{Body: req.Body + " yeah!"}, nil
}
