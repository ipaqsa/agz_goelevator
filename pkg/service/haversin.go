package service

import "math"

const (
	earthR = 6371
)

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func distance(p, q *Point) (m float64) {
	lat1 := degreesToRadians(p.Latitude)
	lon1 := degreesToRadians(p.Longitude)
	lat2 := degreesToRadians(q.Latitude)
	lon2 := degreesToRadians(q.Longitude)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	m = c * earthR * 1000

	return m
}
