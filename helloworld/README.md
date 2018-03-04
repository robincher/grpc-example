gRPC Experiment in Go
======================

Background
-------------

Helloworld Example using gRPC , with both client and server.


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
   $ cd <this-project-repository>/helloword

   $ protoc -I service/  service/helloworld.proto --go_out=plugins=grpc:service
   ```

