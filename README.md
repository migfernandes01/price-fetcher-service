## This small project was made by myself to learn about how to write and strucutre micro-services in Go

This is an HTTP api server that fetches crypto prices implementing a logger and sending metrics.

To run it locally, run `make run`. The server will listen on port `8080` by default.
The ticker (BTC, ETH...) will be extracted from the query params on the `/price` route.
