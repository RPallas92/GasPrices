package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"rpallas/oil/prices/gasstations"
	"rpallas/oil/prices/positions"
	"rpallas/oil/prices/positions/geocoding"
	"rpallas/oil/prices/routes"
	"rpallas/oil/prices/routes/directions"
	"time"
)

func main() {
	gasStationsRepository := gasstations.GasStationsRepository{Datasource: &gasstations.GasStationsDatasource{}}
	directionsRepository := directions.DirectionsRepository{Datasource: &directions.DirectionsDatasource{}}
	positionsRepository := geocoding.PositionsRepository{Datasource: &geocoding.PositionsDatasource{}}

	routesController := routes.RoutesController{
		DirectionsRepository:  &directionsRepository,
		GasStationsRepository: &gasStationsRepository,
		PositionsRepository:   &positionsRepository,
	}
	warmup(&gasStationsRepository)
	scheduleWarmUp(&gasStationsRepository)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/findGasStations/coordinates/:origin/:destination", routesController.FindGasStationsByCoordinates)
	r.GET("/findGasStations/cities/:origin/:destination", routesController.FindGasStationsByCities)
	r.Run()
}

func scheduleWarmUp(gasStationsRepository *gasstations.GasStationsRepository) {
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("0 */4 * * *").Do(func() {
		fmt.Println("Executing scheduled warmup")
		warmup(gasStationsRepository)
	}) // every 4 hours
	s.StartAsync()
}

func warmup(gasStationsRepository *gasstations.GasStationsRepository) {
	fmt.Println("Warming up gas stations")
	somePosition := positions.Position{Lat: 42.241915, Lon: 0.412275}
	gasStationsRepository.GetGasStationsByDistance(somePosition, 100)
	fmt.Println("Warm up finished")
}
