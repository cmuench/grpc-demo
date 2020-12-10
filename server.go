package main

import (
	pb "github.com/cmuench/grpc-demo/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":9000"
)

type server struct {
	pb.UnimplementedStockServer
}

func (s *server) Update(ctx context.Context, in *pb.StockUpdateRequest) (*pb.StockUpdateResponse, error) {
	log.Printf("%+v\n", in)

	return &pb.StockUpdateResponse{Ack: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStockServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
