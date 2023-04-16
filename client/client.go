package client 

import (
  "context"
  "net/http"
  "fmt"
  "encoding/json"

  "github.com/migfernandes01/price-fetcher-service/types"
)

type Client struct {
  endpoint string
}

// New returns a new client with the given endpoint.
func New(endpoint string) *Client {
  return &Client{
    endpoint: endpoint,
  }
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
  // create request url (e.g. http://localhost:8080/price?ticker=BTC)
  endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

  // create http request
  req, err := http.NewRequest("get", endpoint, nil)
  if err != nil {
    return nil, err
  }

  // perform http request
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  // check response status code and return error if not 200
  if res.StatusCode != http.StatusOK {
    // decode error response
    httpErr := map[string]any{}
    err = json.NewDecoder(res.Body).Decode(&httpErr)
    if err != nil {
      return nil, err
    }

    return nil, fmt.Errorf("unexpected status code: %d. Error: %s", res.StatusCode, httpErr["error"])
  }

  // decode response 
  priceResponse := new(types.PriceResponse)
  err = json.NewDecoder(res.Body).Decode(priceResponse)
  if err != nil {
    return nil, err
  }

  return priceResponse, nil
}
