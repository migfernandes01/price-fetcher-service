package main

import (
  "context"
  "fmt"
  "log"
)

func main() {
  service := NewLoggingService(NewMetricService(&priceFetcher{}))

  price, err := service.FetchPrice(context.Background(), "BTC")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(price)
}
