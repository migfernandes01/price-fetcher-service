package main

import (
  "time"
  "context"
  "github.com/sirupsen/logrus"
)

// interface that defines the behavior of the logging service
type loggingService struct {
  next PriceFetcher
}

// constructor for the logging service
func NewLoggingService(next PriceFetcher) PriceFetcher {
  return &loggingService{
    next: next,
  }
}

// implementation of the logging service
// log the method call and then call the next service
func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
  // log data once FetchPrice returns
  defer func(begin time.Time) {
    logrus.WithFields(logrus.Fields{
      "duration" : time.Since(begin),
      "err": err,
      "ticker": ticker,
      "price": price,
    }).Info("method FetchPrice")
  }(time.Now())

  return s.next.FetchPrice(ctx, ticker)
}
