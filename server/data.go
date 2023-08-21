package main

import (
	"context"
	"fmt"
	"weather/m/weatherpr"
)

var WeatherData = map[string]*weatherpr.GetWeatherResponse{
	"Moscow": {
		City:               "Moscow",
		TemperatureCelsius: 25,
		Weather:            "Sunny",
	},
	"London": {
		City:               "London",
		TemperatureCelsius: 18,
		Weather:            "Cloudy",
	},
	"Tokyo": {
		City:               "Tokyo",
		TemperatureCelsius: 30,
		Weather:            "Rainy",
	},
}

func (s *weatherServer) GetWeather(ctx context.Context, req *weatherpr.GetWeatherRequest) (*weatherpr.GetWeatherResponse, error) {
	weather, found := WeatherData[req.City]
	if !found {
		return nil, fmt.Errorf("City not found")
	}

	return weather, nil
}
