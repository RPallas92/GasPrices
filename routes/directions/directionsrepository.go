package directions

import (
	"rpallas/oil/prices/positions"
)

type DirectionsRepository struct {
	Datasource *DirectionsDatasource
}

func (repository *DirectionsRepository) GetDirections(from positions.Position, to positions.Position) ([]positions.Position, error) {
	directionsResponse, err := repository.Datasource.getDirections(from, to)

	if err != nil {
		return []positions.Position{}, err
	}

	var foundPositions []positions.Position

	for _, coordinate := range directionsResponse.Features[0].Geometry.Coordinates {
		foundPositions = append(foundPositions, positions.Position{Lat: coordinate[1], Lon: coordinate[0]})
	}

	return foundPositions, nil
}
