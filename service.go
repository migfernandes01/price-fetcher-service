package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// interface that defines the behavior of the service
type PriceFetcher interface {
  FetchPrice(context.Context, string) (float64, error)
}

// service implementation
type priceFetcher struct {}

type BinanceApiResponse struct {
  Price string  `json:"price"`
  Symbol string `json:"symbol"`
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
  // convert ticket to uppercase
  ticker = strings.ToUpper(ticker)
  price, err := fetchPriceFromApi(ticker)
  if err != nil {
    return 0, fmt.Errorf("error when fetching price from api: %w", err)
  }

  return price, nil
}

func fetchPriceFromApi(ticker string) (float64, error) {
  // build symbol (bincance api format)
  symbol := ticker + "USD"

  // call http api
  res, err := http.Get("https://api.binance.us/api/v3/ticker/price?symbol=" + symbol)
  if err != nil {
    return 0, fmt.Errorf("error when calling binance api: %w", err)
  }

  apiRes := BinanceApiResponse{}
  err = json.NewDecoder(res.Body).Decode(&apiRes)
  if err != nil {
    return 0, fmt.Errorf("error when decoding binance api response: %w", err)
  }

  price, err := strconv.ParseFloat(apiRes.Price, 64)
  if err != nil {
    return 0, fmt.Errorf("error when parsing price from string to float64: %w", err)
  }

  return price, nil
}
