package main

import (
	"buf-tutorial/gen/acme/weather/v1/weatherv1connect"
	"buf-tutorial/handler"
	"buf-tutorial/interceptors"

	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	interceptors := connect.WithInterceptors(interceptors.NewLoggerInterceptor())

	weatherHandler := &handler.WeatherHandler{}
	mux := http.NewServeMux()
	path, handler := weatherv1connect.NewWeatherServiceHandler(weatherHandler, interceptors)

	mux.Handle(path, handler)

	fmt.Printf("Starting server on localhost:8080\n")

	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}))

	fmt.Println(path)
}
