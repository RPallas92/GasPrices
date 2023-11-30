package gasstations

import (
	"io/ioutil"
	"net/http"
	"os"
	"rpallas/oil/prices/gasstations/data"
)

type GasStationsDatasource struct {
}

func (datasource *GasStationsDatasource) LoadGasStations() (data.GasStationsResponse, error) {
	response, err := http.Get("https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/EstacionesTerrestres/")
	if err != nil {
		return data.GasStationsResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data.GasStationsResponse{}, err
	}

	return data.UnmarshalGasStationsResponse(body)
}

func (datasource *GasStationsDatasource) LoadGasStationsFromFile() (data.GasStationsResponse, error) {
	// TODO Ricardo unused - will be used for testing purposes
	fileContent, err := os.ReadFile("prices.json")
	if err != nil {
		return data.GasStationsResponse{}, err
	}
	return data.UnmarshalGasStationsResponse(fileContent)
}
