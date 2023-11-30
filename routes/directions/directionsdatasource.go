package directions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"rpallas/oil/prices/positions"
	"rpallas/oil/prices/routes/directions/data"
)

type DirectionsDatasource struct {
}

func (datasource *DirectionsDatasource) getDirections(from positions.Position, to positions.Position) (data.RouteDirectionsResponse, error) {
	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf("https://api.openrouteservice.org/v2/directions/driving-car?api_key=%v&start=%v,%v&end=%v,%v",
		apiKey, from.Lon, from.Lat, to.Lon, to.Lat)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return data.RouteDirectionsResponse{}, err
	}

	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return data.RouteDirectionsResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data.RouteDirectionsResponse{}, err
	}

	return data.UnmarshalRouteDirectionsResponse(body)
}
