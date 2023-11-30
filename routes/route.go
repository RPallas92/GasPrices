package routes

import (
	"rpallas/oil/prices/positions"
)

type Route struct {
	Origin      positions.Position
	Destination positions.Position
	waypoints   []positions.Position
}
