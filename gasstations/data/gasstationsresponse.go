package data

import (
	"encoding/json"
	"time"
)

func UnmarshalGasStationsResponse(data []byte) (GasStationsResponse, error) {
	var r GasStationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type GasStationsResponse struct {
	Date        string       `json:"Fecha"`
	GasStations []GasStation `json:"ListaEESSPrecio"`
}

func (r *GasStationsResponse) GetTruncatedToDayDate() (time.Time, error) {
	layout := "02/01/2006 15:04:05"
	date, err := time.Parse(layout, r.Date)
	if err != nil {
		return time.Time{}, err
	}

	return date.Truncate(24 * time.Hour), nil
}
