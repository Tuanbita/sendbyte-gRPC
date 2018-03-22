package main

import (
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "example/sendbyte-gRPC"
)

const chunkSize = 64 * 1024 // 64 KiB

type chunkSrv struct {
}

func (chunkSrv) Chicken2(req *pb.Request, srv pb.Greeter_Chicken2Server) error {
	chnk := &pb.Chunk{}

	i := 128*1024*1024
	blob := make([]byte,i ) // 128MiB
	rand.Read(blob)

	for currentByte := 0; currentByte < len(blob); currentByte += chunkSize {
		if currentByte+chunkSize > len(blob) {
			chnk.Chunk = blob[currentByte:len(blob)]
		} else {
			chnk.Chunk = blob[currentByte : currentByte+chunkSize]
		}
		if err := srv.Send(chnk); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":10000")

	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()

	pb.RegisterGreeterServer(g, chunkSrv{})
	log.Println("Serving on localhost:10000")
	log.Fatalln(g.Serve(lis))
}
