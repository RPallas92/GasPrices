package data

import (
	"encoding/json"
	"errors"
	"rpallas/oil/prices/positions"
)

func UnmarshalPositionGeocodingResponse(data []byte) (PositionGeocodingResponse, error) {
	var r PositionGeocodingResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type PositionGeocodingResponse struct {
	Features []Feature `json:"features"`
}

func (p *PositionGeocodingResponse) ToPosition() (positions.Position, error) {
	if p.Features == nil || len(p.Features) == 0 {
		return positions.Position{}, errors.New("Ivalid origin or destination text")

	}
	coordinates := p.Features[0].Geometry.Coordinates
	return positions.Position{Lat: coordinates[1], Lon: coordinates[0]}, nil
}
