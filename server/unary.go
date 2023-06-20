package main

import(
	pb "github.com/muaj07/grpc_demo/proto"
	context "context"
)

/*
	#######---GreetServiceServer interface methods---#########
	type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	SayHelloServerStreaming(*NameList, GreetService_SayHelloServerStreamingServer) error
	SayHelloClientStreaming(GreetService_SayHelloClientStreamingServer) error
	SayHelloBidirectional(GreetService_SayHelloBidirectionalServer) error
	mustEmbedUnimplementedGreetServiceServer()
}
	*/


// Define a function called SayHello that accepts a context and a request as inputs and returns a HelloResponse and an error.
func (m *myGreetingServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	// Create a new HelloResponse object with the message "Hello!" and return it along with nil error.
	return &pb.HelloResponse{Message: "Hello!"}, nil
}