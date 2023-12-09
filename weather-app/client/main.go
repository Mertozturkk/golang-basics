package main

import (
	"context"
	"fmt"
	"io"
	"weather/api"

	"google.golang.org/grpc"
)

func main() {

	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Cities:")
	for _, city := range resp.Items {
		fmt.Printf("%s (%s)\n", city.CityName, city.CityCode)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{CityCode: "ank"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Weather in Ankara: ")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else {
			fmt.Printf("Temperature: %.2f\n", msg.Temerature)
		}
	}
	fmt.Println("server stopped")

}
