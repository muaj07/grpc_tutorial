package main

import (
	"log"
	"time"

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

// SayHelloServerStreaming sends a stream of HelloResponse messages to the client
// based on a list of names received in the request.
func (m *myGreetingServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	/*
					 Stream implements 
		type GreetService_SayHelloServerStreamingServer interface {
    	Send(*HelloResponse) error
    	grpc.ServerStream
}
	*/
	// Print the name received in the request
	log.Printf("Got Request with Name: %v", req.Name)

	// Iterate over the list of names received in the request
	for _, name := range req.Name {
		// Create a HelloResponse message with a personalized greeting
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		} 

		// Send the message to the client
		// Send(*HelloResponse) error
		if err := stream.Send(res); err != nil {
			return err
		}
		// Wait for 2 seconds before sending the next message
		time.Sleep(2 * time.Second)
	}

	// Return nil to indicate success
	return nil
}