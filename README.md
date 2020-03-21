# go-proto-exercise
Go protobuf greeter program


# Contents
1. proto dir:
    Contains all the proto files.

2. SimpleClient / SimpleServer folders using gRPC
    Contains a simple rpc implementation of gRPC based client and Server.

3. StreamingServerClient / StreamingServer folders using gRPC
    Contains a Bi-directional streaming implementation of gRPC based client and Server.

# Tech Stack
- Golang
- gRPC
- Protobuf
- go modules (for dependency management)
- Make
- Docker & Docker Compose

# How to download and execute
1. Git clone to a directory under $GOPATH
2. 'cd' into go-proto-exercise
3. Run 'make' to see the available targets
4. Brief on few commands which one might need to use:
    - 'make pb-codegen': Generates protobuf code.
    - 'make simple-container-build' / 'make stream-container-build': Builds code and constructs container images.
    - 'make simple-start' / 'make stream-start': Starts container instances from previously generated images
    - 'make simple-stops' / 'make stream-stop': Stops container instances
    - 'make simple-config-check' / 'make stream-config-check': Validates docker-compose files

Typical usage ('xxxx' being 'simple' / 'stream'):
- make pb-codegen
- make xxxx-container-build
- make xxxx-start
- make xxxx-stop