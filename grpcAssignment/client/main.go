package main

import (
	"context"
	"fmt"
	"github.com/akshay1192/nomercy/grpcAssignment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"math"
)

const (
	target = "localhost:4040"
	crt = "/Users/akshay/Tekion/goWorkspace/src/github.com/akshay1192/nomercy/grpcAssignment/server.crt"
)

// nums where we will be adding incoming numbers
var nums []int64

func main() {


	// we will use this nums to push in numbers
	go addNumber()

	// Create the client TLS credentials
	cred, err := credentials.NewClientTLSFromFile(crt, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// creating a connection here with grpc server running on port : 4040
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}

	// creating a grpc client in order to call rpc
	client := proto.NewTestServiceClient(conn)

	// getting the stream which will be used as a channel to communicate with client and server (bi directional)
	stream, err := client.GetMax(context.Background())
	if err != nil {
		panic(err)
	}



	// iterating over the nums and sending the values coming from the above addNumber routine
	for _, val := range nums {

		stream.Send(&proto.Request{Num: int64(val)})

		// response from server
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		fmt.Println("Max number after sending ", int64(val), " is : ", in.Result)

	}
}

// func adds new number in the nums slice
func addNumber() {

	// initialising num from 0
	var ctr int64 = 0

	// adding the numbers to nums in a separate go routine
	for {

		if ctr == math.MaxInt64 {
			ctr = 0
		}

		nums = append(nums, ctr)
		ctr++

	}
}