package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"
	"weather/api"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myweatherService{})
	fmt.Println("Server is running on port 8080")
	panic(srv.Serve(lis))
}

type myweatherService struct {
	api.UnimplementedWeatherServiceServer //
}

func (m *myweatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "ank", CityName: "Ankara"},
			&api.CityEntry{CityCode: "ist", CityName: "Istanbul"},
		},
	}, nil
}

func (m *myweatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{Temerature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil

}
