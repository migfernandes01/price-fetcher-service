package main

import (
  "flag"
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
}
