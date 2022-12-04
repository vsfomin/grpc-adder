package adder

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

func (s *GRPCServer) CreateBook(ctx context.Context, req *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	f, err := os.OpenFile("./data/books.txt", os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cat't open file: %w", err)
	}
	defer f.Close()
	book := req.GetBook()
	json, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		return nil, fmt.Errorf("can't marshal json: %w", err)
	}
	if err := os.WriteFile("./data/books.txt", json, 0644); err != nil {
		return nil, fmt.Errorf("can't write json to file: %w", err)
	}
	return &proto.CreateBookResponse{Bid: int64(book.Bid)}, nil
}
