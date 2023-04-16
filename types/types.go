package types 

// this file contains shared types between the API server and the client 

// struct that defines the response type when fetching a price
type PriceResponse struct {
  Ticker string `json:"ticker"`
  Price float64 `json:"price"`
}
