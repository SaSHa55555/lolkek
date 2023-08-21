package main

import (
	"fmt"
	"log"
	"net"
	"weather/m/weatherpr"

	"google.golang.org/grpc"
)

type weatherServer struct {
	weatherpr.UnimplementedOpenWeatherMapServer
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	weatherpr.RegisterOpenWeatherMapServer(server, &weatherServer{})
	fmt.Println("Server is listening on :9090")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
