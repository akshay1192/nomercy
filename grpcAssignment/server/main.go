package main

import (
	"container/heap"
	"fmt"
	"github.com/akshay1192/nomercy/grpcAssignment/proto"
	"github.com/akshay1192/nomercy/grpcAssignment/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	network = "tcp"
	port    = ":4040"
	crt = "/Users/akshay/Tekion/goWorkspace/src/github.com/akshay1192/nomercy/grpcAssignment/server.crt"
	key = "/Users/akshay/Tekion/goWorkspace/src/github.com/akshay1192/nomercy/grpcAssignment/server.key"
)

// server implements all the methods of TestServiceServer interface present in proto generated code
// hence server becomes TestServiceServer type also
type server struct{}

// Business logic for finding the maximum num yet
// Method 1 :
// We can do it by calling sort (in decreasing order by changing the comparator) every time we receive a number
// and then returning the 0th index number

// Method 2 :
// Using priority queue (building a max heap)

func (s *server) GetMax(stream proto.TestService_GetMaxServer) error {

	var pq util.PriorityQueue

	// listening to client and retrieving top element from the max heap
	for {

		req, err := stream.Recv()
		if err != nil {
			fmt.Println("Connection closed from client")
			return err
		}

		num := req.GetNum()
		fmt.Println("Incoming number from client : ", num)

		i := 0

		// building heap / inserting incoming num in the heap
		if len(pq) == 0 {
			pq = append(pq, &util.Item{Priority: num, Index: i})
			i++
			heap.Init(&pq)
		} else {
			heap.Push(&pq, &util.Item{Priority: num, Index: i})
			i++
		}

		//sort.Slice(numbersList,func(i,j int) bool{return numbersList[i] > numbersList[j]})

		// sending the response to client
		res,err := util.GetMaxFromHeap(pq)
		if err != nil  {
			fmt.Printf("error in getting max element from heap : %s" +  err.Error())
			return err
		}

		if err := stream.Send(&proto.Response{Result: res}); err != nil {
			return err
		}
	}
}

func main() {

	listener, err := net.Listen(network, port)
	if err != nil {
		panic(err)
	}

	// Create the TLS credentials
	cred, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(cred)}

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterTestServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
