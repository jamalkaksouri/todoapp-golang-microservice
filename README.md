# todoapp-golang-microservice
 TodoApp-Golang-Microservice

##Installation

# Ensure that MySQL is installed on your computer before proceeding

1. Go to `cd cmd/server`
2. Run command `go build .`
3. Run command `./server.exe -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>`
4. Go to `cd cmd/client-grpc`
5. Run command `go build .`
6. Run command `./client-grpc.exe -server=localhost:9090`