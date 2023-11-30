package data

import "encoding/json"

func UnmarshalRouteDirectionsResponse(data []byte) (RouteDirectionsResponse, error) {
	var r RouteDirectionsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RouteDirectionsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RouteDirectionsResponse struct {
	Features []Feature `json:"features"`
}
