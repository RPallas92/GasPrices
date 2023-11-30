package gasstations

import (
	"errors"
	"fmt"
	"rpallas/oil/prices/datastructures"
	"rpallas/oil/prices/gasstations/data"
	"rpallas/oil/prices/positions"
	"time"
)

type GasStationsRepository struct {
	Datasource              *GasStationsDatasource
	gasStations             *datastructures.KDTree
	gasStationsResponseDate time.Time
}

func (repository *GasStationsRepository) GetGasStationsByDistance(position positions.Position, maxDistance float64) ([]data.GasStation, error) {
	gasStations, err := repository.getAllGasStations()
	if err != nil {
		return []data.GasStation{}, err
	}

	closeGasStationsIndexes := gasStations.Within(position, maxDistance)

	closeGasStations := []data.GasStation{}
	for _, gasStationIndex := range closeGasStationsIndexes {
		gasStation := repository.gasStations.Points[gasStationIndex].(data.GasStation)
		if gasStation.PriceDieselA == "" {
			continue
		}
		closeGasStations = append(closeGasStations, gasStation)
	}

	return closeGasStations, nil
}

func (repository *GasStationsRepository) getAllGasStations() (*datastructures.KDTree, error) {
	if repository.gasStationsResponseDate.IsZero() {
		return repository.getAllGasStationsFromDatasource()
	}

	if today := time.Now().Truncate(24 * time.Hour); today.After(repository.gasStationsResponseDate) {
		return repository.getAllGasStationsFromDatasource()
	}

	return repository.gasStations, nil
}

func (repository *GasStationsRepository) getAllGasStationsFromDatasource() (*datastructures.KDTree, error) {
	fmt.Println("Getting gas stations from datasource")
	gasStationsResponse, err := repository.Datasource.LoadGasStations() // TODO RIcardo
	if err != nil {
		return &datastructures.KDTree{}, err
	}
	if gasStationsResponse.GasStations == nil {
		return &datastructures.KDTree{}, errors.New("nil gas stations")
	}

	date, err := gasStationsResponse.GetTruncatedToDayDate()
	if err != nil {
		return &datastructures.KDTree{}, err
	}

	var points []datastructures.Point
	for _, gasStation := range gasStationsResponse.GasStations {
		price, e := gasStation.GetPriceDieselA()
		if e != nil || price == 0.0 {
			continue
		}
		_, e = gasStation.GetPosition()
		if e != nil {
			continue
		}

		points = append(points, gasStation)
	}
	kdtree := datastructures.NewTree(points, 64)
	repository.gasStations = kdtree
	repository.gasStationsResponseDate = date

	return kdtree, nil
}
