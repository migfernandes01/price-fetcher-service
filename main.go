package main

import (
  "flag"
  //"context"
  //"fmt"

  //"github.com/migfernandes01/price-fetcher-service/client"
)

func main() {
  // extract listenAddr from command line flags
  listenAddr := flag.String("listen-addr", ":8080", "address to listen on")

  // create service instance (with logging and metrics)
  service := NewLoggingService(NewMetricService(&priceFetcher{}))

  // create JSON API server instance
  server := NewJSONAPIServer(*listenAddr, service)

  // start HTTP server
  server.ServeHTTP()

  // use go client to fetch price:
  //client := client.New("http://localhost:8080/price")
  //price, err := client.FetchPrice(context.Background(), "BTC")
  //if err != nil {
    //panic(err)
  //}
  //fmt.Printf("price: %v", price)
  //return
}
