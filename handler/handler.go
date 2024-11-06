package handler

import (
	weatherv1 "buf-tutorial/gen/go/acme/weather/v1"
	"context"

	"connectrpc.com/connect"
)

type WeatherHandler struct{}

func (h *WeatherHandler) GetWeather(cxt context.Context, req *connect.Request[weatherv1.WeatherRequest]) (*connect.Response[weatherv1.WeatherResponse], error) {
	response := &weatherv1.WeatherResponse{
		Temperature: 20,
		Description: "Sunny",
		Location:    req.Msg.Location,
	}

	return connect.NewResponse(response), nil
}
