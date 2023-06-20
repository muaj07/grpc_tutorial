package main

import(
	"log"
	"context"
	"time"
	pb "github.com/muaj07/grpc_demo/proto"
)

/*
	The client is an instant of pb.GreetServiceClient, which will implement the following methods
	type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloServerStreaming(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error)
	SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error)
	SayHelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalClient, error)
}
*/

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.NoParam{} //Argument for the client.SayHello
	// SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	//func (pb.GreetServiceClient).SayHello(ctx context.Context, in *pb.NoParam, opts ...grpc.CallOption) (*pb.HelloResponse, error)
	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message) //res = *pb.HelloResponse
}