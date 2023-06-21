# gRPC Tutorial

This repository contains a simple demonstration of gRPC in Go.

## Dependencies

The project is written in Go and uses the following dependencies:

- google.golang.org/grpc v1.56.0
- google.golang.org/protobuf v1.30.0
- github.com/golang/protobuf v1.5.3
- golang.org/x/net v0.9.0
- golang.org/x/sys v0.7.0
- golang.org/x/text v0.9.0
- google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1

## Makefile

The Makefile contains a command to generate gRPC code from a .proto file:

```makefile
generate_grpc:
	protoc \
	--go_out=. \
	--go-grpc_out=. \
	proto/greet.proto```



## Description of tutorial

This is a very basic tutorial for how to define and implement the 4-type of rpc in Golang using "Protobuf" and "proto compile". Some steps:
#### Step 1: Define the different rpc interfaces in "greet.proto" file inside the "proto2 folder

#### Step 2: Run the "make generate_grpc" command for which you need to install "make"

#### Step 3: Implement all the 4 interfaces define in the "service GreetService" of the "greet.proto" by using the Golang interfaces generated in the "greet_grpc.pb.go" file inside the "proto" folder. 

#### FYI, the "Server" side of the interfaces are defined/implemented in the "Server" folder, while the "Client" side of the interfaces are defined/implemented in the "client" folder. Check the comments for how the methods for different interfaces (created in the "greet_grpc.pb.go") are implemented. 


