package main

import(
	"log"
	"context"
	"time"
	pb "github.com/muaj07/grpc_demo/proto"
	"io"
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


// callSayHelloServerStream is a function that calls the SayHelloServerStreaming 
// method of a pb.GreetServiceClient with the specified context and name list.
func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
    // Create a Context with a timeout of 50 seconds and a function to cancel the Context when it is done.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	
    // Call the SayHelloServerStreaming method with the created context and name list, and get a stream and error in return.
	// Template of the Method taken from "greet_grpc.pb.go"
	// SayHelloServerStreaming(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error)
	stream, err := client.SayHelloServerStreaming(ctx, names) //stream= pb.GreetService_SayHelloServerStreamingClient

	/*
		Stream is an instant of the following interface and will implement its methods

	type GreetService_SayHelloServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}
	*/
	
    // If there is an error in getting the stream, log the error and exit the function.
	if err != nil {
		log.Fatalf("could not send names: %v", stream)
	}
	
    // Loop until there are no more messages in the stream.
	for {
        // Get the next message and error in the stream, and exit the loop if there are no more messages.
		message, err := stream.Recv() //Recv() (*HelloResponse, error)
		if err == io.EOF {
			break
		}
		
        // If there is an error in getting the message from the stream, log the error and exit the function.
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		
        // Log the received message.
		log.Printf("Message received: %s", message.Message)
	}
    // Log that the streaming has finished.
	log.Printf("Streaming Finished")
}