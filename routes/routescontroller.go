package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"rpallas/oil/prices/gasstations"
	"rpallas/oil/prices/positions"
	"rpallas/oil/prices/positions/geocoding"
	"rpallas/oil/prices/routes/directions"
	"strconv"
	"strings"
)

type RoutesController struct {
	GasStationsRepository *gasstations.GasStationsRepository
	DirectionsRepository  *directions.DirectionsRepository
	PositionsRepository   *geocoding.PositionsRepository
}

func (c *RoutesController) FindGasStationsByCoordinates(context *gin.Context) {
	originRaw := context.Param("origin")
	destinationRaw := context.Param("destination")
	origin, destination, err := c.getOriginAndDestination(originRaw, destinationRaw)
	if err != nil {
		context.JSON(400, gin.H{
			"Bad request": err,
		})
	}

	routePlanner := RoutePlanner{
		DirectionsRepositoiry:          c.DirectionsRepository,
		GasStationsRepository:          c.GasStationsRepository,
		PositionsRepository:            c.PositionsRepository,
		GasStationsRouteMaxDistance:    600,
		GasStationsPositionMaxDistance: 8000,
	}

	gasStationsResult, err := routePlanner.FindGasStations(origin, destination)
	if err != nil {
		context.JSON(500, gin.H{
			"Error": "Error getting gas stations",
		})
	}

	context.JSON(200, gasStationsResult)
}

func (c *RoutesController) FindGasStationsByCities(context *gin.Context) {
	origin := context.Param("origin")
	destination := context.Param("destination")

	routePlanner := RoutePlanner{
		DirectionsRepositoiry:          c.DirectionsRepository,
		GasStationsRepository:          c.GasStationsRepository,
		PositionsRepository:            c.PositionsRepository,
		GasStationsRouteMaxDistance:    600,
		GasStationsPositionMaxDistance: 8000,
	}

	gasStationsResult, err := routePlanner.FindGasStationsByCities(origin, destination)
	if err != nil {
		context.JSON(500, gin.H{
			"Error": "Error getting gas stations",
		})
	}

	context.JSON(200, gasStationsResult)
}

func (c *RoutesController) getOriginAndDestination(originRaw, destinationRaw string) (positions.Position, positions.Position, error) {
	origin := strings.Split(originRaw, ",")
	if len(origin) != 2 {
		return positions.Position{}, positions.Position{}, errors.New("Invalid origin coordinates")
	}

	destination := strings.Split(destinationRaw, ",")
	if len(destination) != 2 {
		return positions.Position{}, positions.Position{}, errors.New("Invalid destination coordinates")
	}

	originLat, err := strconv.ParseFloat(origin[0], 64)
	if err != nil {
		return positions.Position{}, positions.Position{}, errors.New("Invalid origin latitude")
	}

	originLon, err := strconv.ParseFloat(origin[1], 64)
	if err != nil {
		return positions.Position{}, positions.Position{}, errors.New("Invalid origin longitude")
	}

	destinationLat, err := strconv.ParseFloat(destination[0], 64)
	if err != nil {
		return positions.Position{}, positions.Position{}, errors.New("Invalid destination latitude")
	}

	destinationLon, err := strconv.ParseFloat(destination[1], 64)
	if err != nil {
		return positions.Position{}, positions.Position{}, errors.New("Invalid destination longitude")
	}

	return positions.Position{Lat: originLat, Lon: originLon}, positions.Position{Lat: destinationLat, Lon: destinationLon}, nil
}
