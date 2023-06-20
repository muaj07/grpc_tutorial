package main

// Import required libraries
import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/muaj07/grpc_demo/proto"
)

const (
	port = ":8089"
)

func main() {
	// Connect to the gRPC server running on the local host
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// If there was an error connecting to the server, log the error and exit
		log.Fatalf("Client cannot connect: %v", err)
	}
	// Close the connection when the function exits
	defer conn.Close()

	// Create a client to call the GreetService on the gRPC server
	//func pb.NewGreetServiceClient(cc grpc.ClientConnInterface) pb.GreetServiceClient
	client := pb.NewGreetServiceClient(conn) //client = pb.GreetServiceClient

	// Create a list of names to send to the GreetService
	names := &pb.NameList{
		Name: []string{"Alice", "Bob", "Eva"},
	}

	// Call the GreetService's SayHelloBidirectionalStream method with the list of names
	callSayHelloBidirectionalStream(client, names)
}
