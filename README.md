
## Description of tutorial

### This is a very basic tutorial for how to define and implement the 4-type of rpc in Golang using "Protobuf" and "proto compile". Some steps:
### Step 1: Define the different rpc interfaces in "greet.proto" file inside the "proto2 folder

### Step 2: Run the "make generate_grpc" command for which you need to install "make"

### Step 3: Implement all the 4 interfaces define in the "service GreetService" of the "greet.proto" by using the Golang interfaces generated in the "greet_grpc.pb.go" file inside the "proto" folder. 

### FYI, the "Server" side of the interfaces are defined/implemented in the "Server" folder, while the "Client" side of the interfaces are defined/implemented in the "client" folder. Check the comments for how the methods for different interfaces (created in the "greet_grpc.pb.go") are implemented. 


