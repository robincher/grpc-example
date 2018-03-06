gRPC Experiment in Go
======================

Background
-------------

Examples to broaden my knowledge in Golang. 

* helloworld : Basic example of a client and server
* grpc-proxy : Example that reads gRPC service definition, and generates a reverse-proxy server which translates a RESTful JSON API into gRPC

Operating ENV
-------------
- Windows 10 / macOS High-Sierra
- Microsoft VS Code + Go Plugins 

Prerequisites
-------------

- This requires Go 1.6 or later
- Requires that [GOPATH is set](https://golang.org/doc/code.html#GOPATH)
- Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

```
$ go help gopath
$ # ensure the PATH contains $GOPATH/bin
$ export PATH=$PATH:$GOPATH/bin
```

Try On it!
-------------

Go to the respective folders to kick start your reference.

References
-------------
* gRPC quickstart (http://www.grpc.io/docs/quickstart/go.html)
* gRPC advance (http://www.grpc.io/docs/tutorials/basic/go.html#generating-client-and-server-code)
* gRPC-gateway a plugin of protoc (https://github.com/grpc-ecosystem/grpc-gateway/)
* Protocol Buffers (https://developers.google.com/protocol-buffers/docs/proto)
* Protoc Explained (https://jbrandhorst.com/post/go-protobuf-tips/)
