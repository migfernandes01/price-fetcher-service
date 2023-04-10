package main

import (
  "context"
  "fmt"
)

// interface that defines the behavior of the service
type PriceFetcher interface {
  FetchPrice(context.Context, string) (float64, error)
}

// service implementation
type priceFetcher struct {}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
  return MockPriceFetcher(ctx, ticker)
}

// mock price for BTC and ETH
var priceMock = map[string]float64{
  "BTC": 20000,
  "ETH": 3000,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
  price, ok := priceMock[ticker]
  if !ok {
    return 0, fmt.Errorf("price not found for %s", ticker)
  }

  return price, nil
}
