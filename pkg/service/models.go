package service

import "sync"

type Point struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type ServiceT struct {
	Source        string
	PointA        *Point
	PointB        *Point
	LongitudeStep int
	LatitudeStep  int
	Area          [][]*Point
	mtx           sync.Mutex
}
type Response struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

type Result struct {
	Dataset   string   `json:"dataset"`
	Elevation float64  `json:"elevation"`
	Location  Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
