proto:
	protoc -I. ./api/proto/v1/*.proto --go-grpc_out=pkg --go_out=pkg

server:
	go run ./cmd/main.go