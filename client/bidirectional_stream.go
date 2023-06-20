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

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	// Print message to indicate that bidirectional streaming has started
	log.Printf("Bidirectional streaming started")

	// Create a context with a timeout of 50 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Establish a bidirectional stream with the server
	//SayHelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalClient, error)
	stream, err := client.SayHelloBidirectional(ctx)//stream=pb.GreetService_SayHelloBidirectionalClient

	/*
		Stream is an instant of and implements the following interface
		type GreetService_SayHelloBidirectionalClient interface {
		Send(*HelloRequest) error
		Recv() (*HelloResponse, error)
		grpc.ClientStream
	}
	*/
	if err != nil {
		log.Fatalf("client could not stream: %v", err)
	}
	// Implement the receiving part of the Stream interface
	// I.e the Recv() (*HelloResponse, error) Method
	// Create a channel to wait for the goroutine to finish
	waitc := make(chan struct{})

	// Start a goroutine to receive messages from the server
	go func() {
		for {
			// Receive a message from the server
			message, err := stream.Recv() //Recv() (*HelloResponse, error)
			// message= *HelloResponse
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			// Print the received message
			log.Println(message.Message)
		}
		// Close the channel when the goroutine is finished
		close(waitc)
	}()

	// Send messages to the server
	//Implement the sending part of the Stream interface
	// i.e implement Send(*HelloRequest) error Method
	for _, name := range names.Name {
		// Create a HelloRequest with the name
		req := pb.HelloRequest{
			Name: name,
		}
		// Send the HelloRequest to the server
		if err := stream.Send(&req); err != nil {
			log.Fatalf("Error while sending request to server %v", err)
		}
		// Print a message indicating that the request has been sent
		log.Printf("Sent the request with name:%v", req.Name)
		// Sleep for 2 seconds between sending each message
		time.Sleep(2 * time.Second)
	}

	// Close the stream for sending messages
	//func (grpc.ClientStream).CloseSend() error
	stream.CloseSend()

	// Wait for the goroutine to finish
	<-waitc

	// Print message to indicate that bidirectional streaming has finished from the client side
	log.Printf("Bidirectional streaming finished from client side")
}