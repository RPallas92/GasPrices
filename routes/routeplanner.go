package routes

import (
	"rpallas/oil/prices/gasstations"
	"rpallas/oil/prices/gasstations/data"
	"rpallas/oil/prices/positions"
	"rpallas/oil/prices/positions/geocoding"
	"rpallas/oil/prices/routes/directions"
	"sort"
)

type RoutePlanner struct {
	DirectionsRepositoiry          *directions.DirectionsRepository
	GasStationsRepository          *gasstations.GasStationsRepository
	PositionsRepository            *geocoding.PositionsRepository
	GasStationsRouteMaxDistance    float64
	GasStationsPositionMaxDistance float64
}

func (r *RoutePlanner) FindGasStations(origin, destination positions.Position) (GasStationsResult, error) {
	route, err := r.calculateRoute(origin, destination)
	if err != nil {
		return GasStationsResult{}, err
	}

	routeGasStations := r.findGasStationsForRoute(route)
	originGasStations := r.findGasStationsForLocation(origin)
	destinationGasStations := r.findGasStationsForLocation(destination)
	routeCoordinates := route.waypoints

	return GasStationsResult{
		RouteGasStations:       routeGasStations,
		OriginGasStations:      originGasStations,
		DestinationGasStations: destinationGasStations,
		RouteCoordinates:       routeCoordinates,
	}, nil
}

func (r *RoutePlanner) FindGasStationsByCities(origin, destination string) (GasStationsResult, error) {
	positions, err := r.PositionsRepository.GetPositions(origin, destination)
	if err != nil {
		return GasStationsResult{}, err
	}

	route, err := r.calculateRoute(positions[0], positions[1])
	if err != nil {
		return GasStationsResult{}, err
	}

	routeGasStations := r.findGasStationsForRoute(route)
	originGasStations := r.findGasStationsForLocation(positions[0])
	destinationGasStations := r.findGasStationsForLocation(positions[1])
	routeCoordinates := route.waypoints

	return GasStationsResult{
		RouteGasStations:       routeGasStations,
		OriginGasStations:      originGasStations,
		DestinationGasStations: destinationGasStations,
		RouteCoordinates:       routeCoordinates,
	}, nil
}

func (r *RoutePlanner) calculateRoute(origin, destination positions.Position) (Route, error) {
	foundDirections, err := r.DirectionsRepositoiry.GetDirections(origin, destination)
	if err != nil {
		return Route{}, err
	}

	return Route{Origin: origin, Destination: destination, waypoints: foundDirections}, nil
}

func (r *RoutePlanner) findGasStationsForRoute(route Route) []data.GasStation {
	var foundGasStations = make(map[string]data.GasStation)

	var lastWaypointSearched positions.Position
	for index, waypoint := range route.waypoints {
		// If next point is too close, don't take it into account as it is redundant
		if index != 0 {
			if distance := waypoint.DistanceFrom(lastWaypointSearched); distance < 450 {
				continue
			}
		}

		waypointGasStations, err :=
			r.GasStationsRepository.GetGasStationsByDistance(waypoint, r.GasStationsRouteMaxDistance)
		if err != nil {
			continue
		}
		lastWaypointSearched = waypoint
		r.addGasStationsToMap(waypointGasStations, foundGasStations)
	}

	routeGasStations := []data.GasStation{}
	for _, value := range foundGasStations {
		routeGasStations = append(routeGasStations, value)
	}
	sort.Sort(data.SortByPrice(routeGasStations))
	return routeGasStations
}

func (r *RoutePlanner) addGasStationsToMap(gasStationsToAdd []data.GasStation, gasStationsMap map[string]data.GasStation) {
	for _, gasStation := range gasStationsToAdd {
		_, exist := gasStationsMap[gasStation.Id]
		if !exist {
			gasStationsMap[gasStation.Id] = gasStation
		}
	}
}

func (r *RoutePlanner) findGasStationsForLocation(position positions.Position) []data.GasStation {
	gasStations, err :=
		r.GasStationsRepository.GetGasStationsByDistance(position, r.GasStationsPositionMaxDistance)
	if err != nil {
		return []data.GasStation{}
	}
	sort.Sort(data.SortByPrice(gasStations))
	return gasStations
}
