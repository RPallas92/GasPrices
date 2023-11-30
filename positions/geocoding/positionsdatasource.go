package geocoding

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"rpallas/oil/prices/positions/geocoding/data"
)

type PositionsDatasource struct {
}

func (datasource *PositionsDatasource) getPosition(city string) (data.PositionGeocodingResponse, error) {
	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%v&text=%v&boundary.country=ES",
		apiKey, url2.QueryEscape(city))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return data.PositionGeocodingResponse{}, err
	}

	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return data.PositionGeocodingResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data.PositionGeocodingResponse{}, err
	}

	return data.UnmarshalPositionGeocodingResponse(body)
}
