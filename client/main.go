package main

import (
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "example/sendbyte-gRPC"
	"fmt"
	"math/rand"
	//"github.com/constabulary/gb/testdata/src/c"
	"strconv"
)
const chunkSize = 64 * 1024 // 64 KiB//= 64*1024 *


func chicken3(client pb.GreeterClient){

	fmt.Println("chicken3: ")
	stream, err := client.Chicken3(context.Background())
	if err != nil {
		log.Fatal("client Chat get stream error: ", err)
	}

	width := 200
	height := 300
	w := []byte(strconv.Itoa(width))
	h := []byte(strconv.Itoa(height))

//1 MiB = 1024*1024 byte
	chnk := &pb.Chunk{}


	t := 1024*1024
	i := 128*t// i = 128 MiB =
	blob := make([]byte,i ) // 128MiB
	rand.Read(blob)
	dem :=0

	for currentByte := 0; currentByte < len(blob); currentByte += chunkSize {

		if dem ==0 {
			chnk.Chunk = w
			stream.Send(chnk)
		} else
		if dem ==1 {
			chnk.Chunk = h
			stream.Send(chnk)
		}else
		if currentByte+chunkSize > len(blob) {
			chnk.Chunk = blob[currentByte:len(blob)]
		} else {
			chnk.Chunk = blob[currentByte : currentByte+chunkSize]
		}
		if err := stream.Send(chnk); err != nil {
		}
		dem++
	}
	fmt.Println("chicken3: ", dem)
}

func chicken4(client pb.GreeterClient){

	fmt.Println("Chat Start:")
	stream, err := client.Chicken4(context.Background())
	if err != nil {
		log.Fatal("client Chat get stream error: ", err)
	}
	chnk := &pb.Chunk{}

	i := 128*1024*1024
	blob := make([]byte,i ) // 128MiB
	rand.Read(blob)

	var blob2 []byte

	for currentByte := 0; currentByte < len(blob); currentByte += chunkSize {

		c, _:= stream.Recv()
		blob2 = append(blob2, c.Chunk...)

		if currentByte+chunkSize > len(blob) {
			chnk.Chunk = blob[currentByte:len(blob)]
		} else {
			chnk.Chunk = blob[currentByte : currentByte+chunkSize]
		}
		if err := stream.Send(chnk); err != nil {
		}
	}
}


func main() {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc := pb.NewGreeterClient(conn)

	chicken3(cc)
/*
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
*/
}