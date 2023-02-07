package service

import (
	"errors"
	"strconv"
	"strings"
)

func parsePoint(data string) (*Point, error) {
	var splits []string
	splits = strings.Split(data, ",")
	if len(splits) != 2 {
		splits = strings.Split(data, " ")
		if len(splits) != 2 {
			return nil, errors.New("wrong format input")
		}
	}
	pLong := strings.TrimSpace(splits[1])
	pLat := strings.TrimSpace(splits[0])
	Long, err := strconv.ParseFloat(pLong, 64)
	if err != nil {
		return nil, err
	}
	Lat, err := strconv.ParseFloat(pLat, 64)
	if err != nil {
		return nil, err
	}
	return newPoint(Lat, Long, -1.0), nil
}

func parseStep(data string) (int, error) {
	return strconv.Atoi(data)
}

func (service *ServiceT) chunk(size int) [][]*Point {
	l := 0
	for _, v := range service.Area {
		l += len(v)
	}
	result := make([]*Point, 0, l)
	for _, v := range service.Area {
		for _, e := range v {
			result = append(result, e)
		}
	}
	chunksSize := l/size + 1
	chunks := make([][]*Point, chunksSize)
	for i := 0; i < chunksSize; i++ {
		if len(result[size*i:]) >= size {
			chunks[i] = result[size*i : size*(i+1)]
		} else {
			chunks[i] = result[size*i:]
		}
	}
	return chunks
}
