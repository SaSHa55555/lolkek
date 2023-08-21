package main

import (
	"context"
	"fmt"
	"log"

	//"net/http"
	"weather/m/weatherpr"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := weatherpr.NewOpenWeatherMapClient(conn)

	var city string
	fmt.Print("Enter city name: ")
	fmt.Scanln(&city)

	req := &weatherpr.GetWeatherRequest{City: city} // Запрос на погоду в указанном городе
	resp, err := client.GetWeather(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("City: %s\nTemperature: %d°C\nWeather: %s\n", resp.GetCity(), resp.GetTemperatureCelsius(), resp.GetWeather())

}
