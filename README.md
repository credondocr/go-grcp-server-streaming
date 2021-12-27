# gRPC Server-Side Streaming with Golang

>A server-streaming RPC is similar to a unary RPC, except that the server returns a stream of messages in response to a client’s request. After sending all its messages, the server’s status details (status code and optional status message) and optional trailing metadata are sent to the client. This completes processing on the server side. The client completes once it has all the server’s messages.

<br/>

# Overview

![](https://grpc.io/img/landing-2.svg)
<br/>

# How to install protoc

Please read this documentation for more [information](http://google.github.io/proto-lens/installing-protoc.html)

<br/>

# How to generate the proto files

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/data.proto`

# Environment Variables

Copy the file `.env.example` into streaming-server directory and remove the example extension
