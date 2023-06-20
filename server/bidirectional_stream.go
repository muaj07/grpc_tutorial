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


func (m *myGreetingServer) SayHelloBidirectional(stream pb.GreetService_SayHelloBidirectionalServer) error {

	/* 
		type GreetService_SayHelloBidirectionalServer interface {
    	Send(*HelloResponse) error
    	Recv() (*HelloRequest, error)
    	grpc.ServerStream
		}
	*/

	// Log server start message
	log.Printf("Server Start Receiving Bidirectional stream")

	// Loop to continuously receive requests from client
	for {
		// Receive request from client
		req, err := stream.Recv() //Recv() (*HelloRequest, error)

		// Note the following struct of HelloRequest
		// message HelloRequest{
		// 	string name = 1;
		// }

		// If end of stream, return
		if err == io.EOF {
			return nil
		}

		// If error receiving stream, log error and return
		if err != nil {
			log.Fatalf("Error while receiving stream from client: %v", err)
			return err
		}

		// Log received request message
		log.Printf("Request from client received: %s", req.Name) //string type

		// Create response message
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		// Note the HelloResponse structure
		//message HelloResponse{
		//	string message = 1;
		// }

		// Send response message to client
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
