## Small project to demonstrate how to write and strucutre micro-services in Go

This project implements both an HTTP server and a gRPC server to fetch crypto prices. This is an example of how to strucutre a go service
by separating concerns and implementing

To run it locally, run `make run`.
To generate proto files, run `make proto`.
The HTTP server will listen on port `8080` by default.
The gRPC server will listen on port `8000` by default.
The ticker (BTC, ETH...) will be extracted from the query params on the `/price` route (GET).

### Folder structure:

- `logging.go`: middleware used for logging
- `metrics.go`: middleware used for metrics
- `service.go`: where the mocked price fetching service is defined
- `json_api.go`: where the json api server is defined
- `grpc_server.go`: where the grpc server is defined
- `/types`: where all the shared types are defined
- `/client`: where the clients (http and grpc) are defined
- `/proto`: where the proto file and go generated grpc files are

### TODO:

[] Use 3rd party service to fetch real crypto prices
[] Dockerize service
