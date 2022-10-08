# TodoApp-golang-microservice
Installation

$\textcolor{orange}{Ensure\ that\ MySQL\ is\ installed\ on\ your\ computer\ before\ proceeding}$

1. Go to directory `cd cmd/server`
2. Run `go build .`
3. Run `./server.exe -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>`
4. Go to `cd cmd/client-grpc`
5. Run `go build .`
6. Run `./client-grpc.exe -server=localhost:9090`

```diff
\textcolor{green}{Or running both servers and clients from a single directory named run apps. (the first server to be run)}
```