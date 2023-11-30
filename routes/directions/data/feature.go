package data

type Feature struct {
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}
