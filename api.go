package main

import (
	"context"
	"encoding/json"

	"math/rand"
	"net/http"

	"github.com/migfernandes01/price-fetcher-service/types"
)

// type of an API function
type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error 

// interface that defines the behavior of the API server
type JSONAPIServer struct {
  listenAddr  string 
  service     PriceFetcher
}

// constructor for the API server
func NewJSONAPIServer(listenAddr string, service PriceFetcher) *JSONAPIServer {
  return &JSONAPIServer{
    listenAddr: listenAddr,
    service: service,
  }
}

// function that creates a new API handler function
func makeHTTPHandlerFunc(fn APIFunc) http.HandlerFunc {
  // create context and add requestID to it
  ctx := context.Background()
  ctx = context.WithValue(ctx, "requestID", rand.Intn(100000000000))
  
  return func(w http.ResponseWriter, r *http.Request) {
    if err := fn(ctx, w, r); err != nil {
      // handle error
      writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
    }
  }
}

// initialize http server
func (s *JSONAPIServer) ServeHTTP() {
  http.HandleFunc("/price", makeHTTPHandlerFunc(s.handleFetchPrice))
  http.ListenAndServe(s.listenAddr, nil)
}

// handler to fetch price and return header/return JSON response
func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
  // get ticker from request
  ticker := r.URL.Query().Get("ticker")

  price, err := s.service.FetchPrice(ctx, ticker)
  if err != nil {
    return err
  }
 
  // create response of type PriceResponse
  priceResponse := types.PriceResponse{
    Ticker: ticker,
    Price: price,
  }

  return writeJSON(w, http.StatusOK, priceResponse)
}

// helper function to write json response
func writeJSON(w http.ResponseWriter, status int, v any) error {
  // write response header with it's status and type
  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")

  // return json encoded response
  return json.NewEncoder(w).Encode(v)
}
