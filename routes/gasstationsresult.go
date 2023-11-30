package routes

import (
	"rpallas/oil/prices/gasstations/data"
	"rpallas/oil/prices/positions"
)

type GasStationsResult struct {
	RouteGasStations       []data.GasStation
	OriginGasStations      []data.GasStation
	DestinationGasStations []data.GasStation
	RouteCoordinates       []positions.Position
}
