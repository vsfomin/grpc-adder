package main

import (
	"log"
	"net"

	"github.com/vsfomin/grpc-adder/api/proto"
	"github.com/vsfomin/grpc-adder/pkg/adder"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	proto.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
