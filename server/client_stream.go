package main

import(
	"log"
	"io"
	pb "github.com/muaj07/grpc_demo/proto"
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

func(m *myGreetingServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	/* Stream is an implementation of 
	type GreetService_SayHelloClientStreamingServer interface {
    SendAndClose(*MessageList) error
    Recv() (*HelloRequest, error)
    grpc.ServerStream
}
	*/

	log.Printf("Start Receiving Streaming")
	var messages []string
	for{
		req, err := stream.Recv() //Recv() (*HelloRequest, error)
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Message: messages}) //SendAndClose(*MessageList) error
		}
		if err !=nil{
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Printf("Request from client received: %s", req.Name)
		messages= append(messages, "Hello " + req.Name)

	}
}
