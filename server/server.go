package main

import (
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "example/sendbyte-gRPC"
	"fmt"

	"io"
	"strconv"
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
func (chunkSrv) Chicken3(stream pb.Greeter_Chicken3Server) error {

	i := 128*1024*1024
	blob := make([]byte,i ) // 128MiB
	rand.Read(blob)
	dem :=0
	for {
		if dem ==0{
			c, _ := stream.Recv()

			value, _ := strconv.ParseInt(string(c.GetChunk()), 10, 64)
			fmt.Println("width: ",value)
			dem ++
		} else
		if dem ==1{
			c, _ := stream.Recv()
			value, _ := strconv.ParseInt(string(c.GetChunk()), 10, 64)
			fmt.Println("height: ",value)
			dem ++
		} else {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful chicken3", len(blob))
				return nil
			}
			//panic(err)
		}
		blob = append(blob, c.GetChunk()...)
		dem ++}
	}


}
func (chunkSrv) Chicken4(stream pb.Greeter_Chicken4Server) error {

	fmt.Println("Begin Chat")
	chnk := &pb.Chunk{}

	i := 128*1024*1024
	blob := make([]byte,i ) // 128MiB
	rand.Read(blob)
	var blob2 []byte

	for currentByte := 0; currentByte < len(blob); currentByte += chunkSize {
		if currentByte+chunkSize > len(blob) {
			chnk.Chunk = blob[currentByte:len(blob)]
		} else {
			chnk.Chunk = blob[currentByte : currentByte+chunkSize]
		}
		if err := stream.Send(chnk); err != nil {
			return err
		}
		c, _:= stream.Recv()

		blob2 = append(blob2, c.Chunk...)
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
