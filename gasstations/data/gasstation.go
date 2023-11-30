package data

import (
	"encoding/json"
	"rpallas/oil/prices/positions"
	"strconv"
	"strings"
)

type GasStation struct {
	PostalCode   string `json:"C.P."`
	Address      string `json:"Dirección"`
	Schedule     string `json:"Horario"`
	Lat          string `json:"Latitud"`
	City         string `json:"Localidad"`
	Lon          string `json:"Longitud (WGS84)"`
	Municipality string `json:"Municipio"`
	PriceDieselA string `json:"Precio Gasoleo A"`
	Province     string `json:"Provincia"`
	Name         string `json:"Rótulo"`
	Id           string `json:"IDEESS"`
}

func (gasStation GasStation) GetPosition() (positions.Position, error) {
	gasStationLat, err := gasStation.getLat()
	if err != nil {
		return positions.Position{}, err
	}
	gasStationLon, err := gasStation.getLon()
	if err != nil {
		return positions.Position{}, err
	}

	return positions.Position{Lat: gasStationLat, Lon: gasStationLon}, nil
}

func (gasStation GasStation) GetPriceDieselA() (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(gasStation.PriceDieselA, ",", "."), 64)
}

func (gasStation GasStation) Coordinates() (float64, float64) {
	position, err := gasStation.GetPosition()
	if err != nil {
		panic("Gas station without position")
	}
	return position.Lon, position.Lat
}

func (gasStation GasStation) Distance(x, y float64) float64 {
	position, err := gasStation.GetPosition()
	if err != nil {
		panic("Gas station without position")
	}
	return position.Distance(x, y)
}

func (gasStation GasStation) getLat() (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(gasStation.Lat, ",", "."), 64)
}

func (gasStation GasStation) getLon() (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(gasStation.Lon, ",", "."), 64)
}

func (g GasStation) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	m["postalCode"] = g.PostalCode
	m["address"] = g.Address
	m["schedule"] = g.Schedule
	m["city"] = g.City
	m["municipality"] = g.Municipality
	m["province"] = g.Province
	m["name"] = g.Name
	m["id"] = g.Id

	lat, err := g.getLat()
	if err == nil {
		m["lat"] = lat
	}
	lon, err := g.getLon()
	if err == nil {
		m["lon"] = lon
	}
	price, err := g.GetPriceDieselA()
	if err == nil {
		m["priceDieselA"] = price
	}

	return json.Marshal(m)
}

// Sorting
type SortByPrice []GasStation

func (s SortByPrice) Len() int {
	return len(s)
}
func (s SortByPrice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortByPrice) Less(i, j int) bool {
	priceI, err := strconv.ParseFloat(strings.ReplaceAll(s[i].PriceDieselA, ",", "."), 64)
	if err != nil {
		return false
	}
	priceJ, err := strconv.ParseFloat(strings.ReplaceAll(s[j].PriceDieselA, ",", "."), 64)
	if err != nil {
		return true
	}
	return priceI < priceJ
}
