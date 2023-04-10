package main

import (
  "context"
  "fmt"
)

// interface that defines the behavior of the metrics service 
type metricService struct { 
  next PriceFetcher
}

// constructor for the metrics service
func NewMetricService(next PriceFetcher) PriceFetcher {
  return &metricService{
    next: next,
  }
}

// implementation of the metrics service 
// send metrics to Datadog/Prometheus and then call the next service
func (s *metricService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
  fmt.Println("Sending metrics to Datadog/Prometheus")
  
  // send metrics to Datadog/Prometheus 

  return s.next.FetchPrice(ctx, ticker)
}
