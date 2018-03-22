package main

import (
	"log"
	"io"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "example/sendbyte-gRPC"
)

func main() {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc := pb.NewGreeterClient(conn)

	client, err := cc.Chicken2(context.Background(), &pb.Request{Req:"hi"})
	if err != nil {
		panic(err)
	}
	var blob []byte
	for {
		c, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				return
			}

			panic(err)
		}

		blob = append(blob, c.Chunk...)
	}
}