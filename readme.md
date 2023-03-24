# Mock gRPC server for client testing 
## Instruction
1. run 
```
$ protoc -I proto/. --go-grpc_out proto/.  --go_out proto/. --go_opt paths=source_relative  proto/example.proto
```
2. mock
```shell
$ mockery --dir=. --inpackage --name=ExampleServiceServer
```
3. execute client_test.go