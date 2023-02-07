package service

var Service *ServiceT

func CreateService() {
	Service = &ServiceT{
		PointA: nil,
		PointB: nil,
		Area:   make([][]*Point, 0),
	}
}

func newPoint(lat, long, alt float64) *Point {
	return &Point{
		Latitude:  lat,
		Longitude: long,
		Altitude:  alt,
	}
}
