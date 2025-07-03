mkdir pb
protoc -I . --go_out=pb --go-grpc_out=pb *.proto