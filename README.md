gRPC Experiment in Go
======================

Background
-------------

Helloworld Example using gRPC. Codes are taken from the official site and meant to further my understanding.

Operating ENV
-------------
- Windows 10 / macOS High-Sierra
- Microsoft VS Code + Go Plugins 

Prerequisites
-------------

- This requires Go 1.6 or later
- Requires that [GOPATH is set](https://golang.org/doc/code.html#GOPATH)

```
$ go help gopath
$ # ensure the PATH contains $GOPATH/bin
$ export PATH=$PATH:$GOPATH/bin
```

Try On it!
-------------

- Run the server

  ```
  $ go run server/main.go
  ```

- Run the client (Open a new terminal)

  ```
  $ go run client/main.go
  ```

Rebuilding the generated code
----------------------------------------

Rebuilding is required whenever you update any service definitions. 

1. Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

2. Install the protoc Go plugin

   ```
   $ go get -u github.com/golang/protobuf/protoc-gen-go
   ```

3. Rebuild the generated Go code

   ```
   $ go generate github.com/username/grpc-example/...
   ```
   
   Or run `protoc` command (with the grpc plugin)
   
   ```
   $ cd <this-project-repository>

   $ protoc -I service/ service/helloworld.proto --go_out=plugins=grpc:service
   ```

References
-------------
* gRPC quickstart (http://www.grpc.io/docs/quickstart/go.html)
* gRPC advance (http://www.grpc.io/docs/tutorials/basic/go.html#generating-client-and-server-code)
* grpc-gateway a plugin of protoc (https://github.com/grpc-ecosystem/grpc-gateway/)
* Protocol Buffers (https://developers.google.com/protocol-buffers/docs/proto)