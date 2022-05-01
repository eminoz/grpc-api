package main

import (
	"context"
	api "github.com/eminoz/grpc-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
}

func main() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)

	}
	newServer := grpc.NewServer()
	api.RegisterAddServiceServer(newServer, &server{})
	reflection.Register(newServer)
	err = newServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *api.Request) (*api.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &api.Response{Result: result}, nil

}
func (s *server) Multiply(ctx context.Context, requst *api.Request) (*api.Response, error) {
	a, b := requst.GetA(), requst.GetB()
	result := a * b
	return &api.Response{Result: result}, nil
}
