package main

import (
  "flag"
  "context"
  "fmt"

  "github.com/migfernandes01/price-fetcher-service/client"
  "github.com/migfernandes01/price-fetcher-service/proto"
)

func main() {
  // extract listen addresses from flags to start HTTP and GRPC servers
  var (
    jsonListenAddr = flag.String("json-listen-addr", ":8080", "address http server will listen on")
    grpcListenAddr = flag.String("grpc-listen-addr", ":8000", "address grpc server will listen on")
  )

  flag.Parse()

  // create service instance (with logging and metrics)
  service := NewLoggingService(NewMetricService(&priceFetcher{}))

  // create and run GRPC server instance in a new routine (since it blocks the execution)
  go makeGRPCServerAndRun(*grpcListenAddr, service)

  // use grpc client to fetch price
  grpcClient, err := client.NewGRPCClient(*grpcListenAddr)
  if err != nil {
    panic(err)
  }

  go func() {
    res, err := grpcClient.FetchPrice(context.Background(), &proto.PriceRequest{Ticker: "BTC"})
    if err != nil {
      panic(err)
    }
    fmt.Printf("price: %v", res)
  }()

  // create JSON API server instance
  jsonServer := NewJSONAPIServer(*jsonListenAddr, service)

  // start HTTP server
  jsonServer.ServeHTTP()



  // use go client to fetch price:
  //client := client.New("http://localhost:8080/price")
  //price, err := client.FetchPrice(context.Background(), "BTC")
  //if err != nil {
  //panic(err)
  //}
  //fmt.Printf("price: %v", price)
  //return
}
