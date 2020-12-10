package main

import (
	"context"
	"log"
	"time"

	pb "github.com/cmuench/grpc-demo/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStockClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Update(ctx, &pb.StockUpdateRequest{
		SourceCode: "default",
		Sku:        "24-MB01",
		Qty:        99.0,
	})
	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Printf("Update ACK: %s", r.Ack)
}
