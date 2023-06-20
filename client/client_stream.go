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


func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	// Print a log message to indicate that client streaming has started.
	log.Printf("Client streaming started")

	// Create a context with a timeout of 50 seconds and defer cancelling it.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Get a stream from the client to send hello requests.
	// SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error)
	stream, err := client.SayHelloClientStreaming(ctx) //stream:= GreetService_SayHelloClientStreamingClient
	/*
		Stream is an instant of the following interface

		type GreetService_SayHelloClientStreamingClient interface {
		Send(*HelloRequest) error
		CloseAndRecv() (*MessageList, error)
		grpc.ClientStream
}
	*/

	if err != nil {
		// If there is an error, log the error and exit.
		log.Fatalf("client could not stream: %v", stream)
	}

	// Loop through each name in the name list.
	for _, name := range names.Name {
		// Create a new request with the current name.
		// This will use an argument for the Send Method
		// of the Stream
		req := &pb.HelloRequest{
			Name: name,
		}

		// Send the request to the server.
		if err := stream.Send(req); err != nil {
			// If there is an error sending the request, log the error and exit.
			log.Fatalf("Error while sending req:%v", req)
		}

		// Print a log message to indicate that the request has been sent.
		log.Printf("Sent the request with name:%v", req.Name)

		// Wait for 2 seconds before sending the next request.
		time.Sleep(2 * time.Second)
	}

	// Close the stream and receive the server response.
	// This is the implementation of the following method
	// CloseAndRecv() (*MessageList, error)

	res, err := stream.CloseAndRecv() //res=*MessageList
	if err != nil {
		// If there is an error receiving the response, log the error and exit.
		log.Fatalf("Error while receiving response from Server %v", err)
	}

	// Print the server response.
	log.Printf("Response from the Server %v", res.Message)

	// Print a log message to indicate that the client streaming has finished.
	log.Printf("Streaming Finished")
}