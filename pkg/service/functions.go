package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func (service *ServiceT) GetAltitudes(a, b, longstep, latstep string, source string) error {
	service.mtx.Lock()
	defer service.mtx.Unlock()
	point, err := parsePoint(a)
	if err != nil {
		return err
	}
	service.PointA = point

	point, err = parsePoint(b)
	if err != nil {
		return err
	}
	service.PointB = point

	step, err := parseStep(latstep)
	if err != nil {
		return err
	}
	service.LatitudeStep = step

	step, err = parseStep(longstep)
	if err != nil {
		return err
	}
	service.LongitudeStep = step

	service.Source = source
	service.generate()
	chunks := service.chunk(100)

	wg := sync.WaitGroup{}
	wg.Add(len(chunks))
	for i, chunk := range chunks {
		if len(chunk) == 0 {
			wg.Done()
			break
		}
		err = service.altitude(&wg, chunk, 0)
		if err != nil {
			return err
		}
		fmt.Printf("%d/%d\n", i*100, len(chunks)*100)
	}
	wg.Wait()
	println("Success")
	err = service.writeExcel("data/data.xlsx")
	if err != nil {
		return err
	}
	println("Done")
	return nil
}

func (service *ServiceT) altitude(wg *sync.WaitGroup, points []*Point, r int) error {
	url := fmt.Sprintf("https://api.opentopodata.org/v1/%s?locations=", service.Source)
	for _, point := range points {
		url += fmt.Sprintf("%f,%f|", point.Latitude, point.Longitude)
	}
	url = url[:len(url)-1]
	get, err := http.Get(url)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(get.Body)
	var resp Response
	err = decoder.Decode(&resp)
	if err != nil {
		return err
	}
	if len(resp.Results) == 0 {
		if r == 50 {
			return errors.New(resp.Status)
		}
		time.Sleep(time.Second * 1)
		return service.altitude(wg, points, r+1)
	}
	for i, point := range points {
		point.Altitude = resp.Results[i].Elevation
	}
	wg.Done()
	return nil
}
