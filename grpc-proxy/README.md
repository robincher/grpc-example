gRPC Experiment in Go
======================

Background
-------------

gPRC example that depicts how to define a serivce in a .proto file, generate service and client code using protoc and deplying a simple gRPC api that is consumable using HTTP REST call.

![grpc-gateway-overview.png](https://github.com/robincher/grpc-example/blob/master/assets/grpc-gateway-overview.png)


Try On it!
-------------

- Run the server

```
$ cd grpc-proxy
$ go run server/echo_server.go
```

- Run the client (Open a new terminal)

```
$ go run client/echo_client.go
```

- Run the Proxy 

```
$ go run proxy/echo_proxy.go
```

- Do a test API call

```
curl -X POST "http://localhost:8080/v1/example/echo/1/username"
```

You should get a response 

```
{
"id": "1",
"msg": "username"
}
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
   $cd <this-project-repository>/grpc-proxy
    
   # Generate Client and Server Stub
   $ protoc -I. -I %GOPATH%/src -I  service/ service/echo.proto   -I %GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:.

   # Generate Proxy Stub 
   $ protoc -I. -I %GOPATH%/src -I service/  service/echo.proto -I %GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.
   ```

