package main

import (
  "context"
  "net"
  "math/rand"

  "github.com/migfernandes01/price-fetcher-service/proto"
  "google.golang.org/grpc"
)

// create and run GRPC server 
func makeGRPCServerAndRun(listenAddr string, service PriceFetcher) error {
  // new grpc price fetcher passing the service implementation
  grpcPriceFetcherServer := NewGRPCPriceFetcher(service)

  // listen on the specified address
  ln, err := net.Listen("tcp", listenAddr)
  if err != nil {
    return err
  }

  // create new grpc server
  opts := []grpc.ServerOption{}
  server := grpc.NewServer(opts...)

  // register the grpc server with the grpc server and the grpc price fetcher service
  proto.RegisterPriceFetcherServer(server, grpcPriceFetcherServer)

  // start the grpc server
  return server.Serve(ln)
}

// define the interface that the server will implement
// we need to implement the grpc interface (see it in service_grpc.pb.go)
type GRPCPriceFetcherServer struct {
  service PriceFetcher
  proto.UnimplementedPriceFetcherServer // used by grpc 
}

// constructor for the grpc server
func NewGRPCPriceFetcher(service PriceFetcher) *GRPCPriceFetcherServer {
  return &GRPCPriceFetcherServer{
    service: service,
  }
}

// implementation of the FetchPrice method of the GRPCPriceFetcherServer
// using generated types from the proto file
func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
  // add a requestId to context
  ctx = context.WithValue(ctx, "requestID", rand.Intn(1000))
  // call the service implementation to fetch the price
  price, err := s.service.FetchPrice(ctx, req.Ticker)
  if err != nil {
    return nil, err
  }

  res := &proto.PriceResponse{
    Ticker: req.Ticker,
    Price: float32(price),
  }

  return res, nil 
}
