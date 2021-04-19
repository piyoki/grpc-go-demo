# gRPC Go Demo

## gRPC Introduction

Remote Procedure Calls are something that we use within distributed systems that allow us to communicate between applications. More specifically, it allows us to expose methods within our application that we want other applications to be able to invoke.

It’s similar to REST API communication in the sense that with it, you are effectively exposing functionality within your app to other apps using a HTTP connection as the communication medium.

## Differences between gRPC and REST

Whilst REST and gRPC are somewhat similar, there are some fundamental differences in how they work that you should be aware of.

1. gRPC utilizes `HTTP/2` whereas REST utilizes `HTTP 1.1`
2. gRPC utilizes the protocol buffer data format as opposed to the standard JSON data format that is typically used within REST APIs
3. With gRPC you can utilize `HTTP/2` capabilities such as server-side streaming, client-side streaming or even bidirectional-streaming should you wish.

## Prerequisites

You may need to have the following installed on your local machine:

```bash
# Install protobuff compiler
$ sudo apt install protobuf-compiler
# Check Version
$ protoc --version  # Ensure compiler version is 3+
# Install protoc-gen-go
$ go get -u google.golang.org/grpc # install grpc library
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

You will have to ensure that `$GOPATH/bin` is on your environment path so that you can use the `protoc` tool later on in this tutorial.
Add the following lines to your `.bashrc` or `.zshrc` file

```bash
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH/bin
```

## Run the Demo

Spins up the server

```bash
go run server.go
```

Initiate a client connection and send a message to the server

```bash
go run client/client.go
```

## Reference

[gRPC Offical Documentation](https://grpc.io/docs/languages/go/quickstart/)
