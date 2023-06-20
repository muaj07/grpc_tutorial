package main

import(
	"log"
	"net"
	grpc "google.golang.org/grpc"
	pb "github.com/muaj07/grpc_demo/proto"
)

const(
	port= ":8089"
)

// Define a struct named myGreetingServer
type myGreetingServer struct {
    // Embed the GreetServiceServer interface from the pb package into this struct
	//By embedding pb.GreetServiceServer into our myGreetingServer struct, 
	//we gain access to all of the methods and fields defined in the pb.GreetServiceServer 
	//interface. This is a way of achieving composition in Go that allows us to create
	// new types by combining existing types.
    pb.GreetServiceServer

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
}


func main(){
// net package provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
// The net.Listen function creates a TCP listener on the given address.
lis, err := net.Listen("tcp", port)
if err !=nil {
    // If there's an error, log the error and exit with a fatal error.
    log.Fatalf("failed to listen: %v", err)
}

// grpc.NewServer creates a new gRPC server instance.
grpcServer := grpc.NewServer()

// Instantiate our custom server.
service := &myGreetingServer{}

// Register the service implementation with the gRPC server.
// pb.RegisterGreetServiceServer is a generated function created by proto compile.
// func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
// 	s.RegisterService(&GreetService_ServiceDesc, srv)
// }

pb.RegisterGreetServiceServer(grpcServer, service)



// Print the listener's address to the logs.
log.Printf("Serve started at %v", lis.Addr())

// Start the gRPC server.
err = grpcServer.Serve(lis)
if err !=nil{	
    // If there's an error, log the error and exit with a fatal error.
    log.Fatalf("cannot start server: %v", err)
}
}