package service

func getSteps(point *Point, longstep, latstep int) (float64, float64) {
	var dis float64

	longstepT := 0.0
	count := 0
	for count < 1000000 {
		tmp := newPoint(point.Latitude, point.Longitude+longstepT, 0)
		dis = distance(point, tmp)
		if dis >= float64(longstep) {
			break
		}
		count++
		longstepT += 0.0000001
	}

	latstepT := 0.0
	count = 0
	for count < 1000000 {
		tmp := newPoint(point.Latitude+latstepT, point.Longitude, 0)
		dis = distance(point, tmp)
		if dis >= float64(latstep) {
			break
		}
		count++
		latstepT += 0.0000001
	}
	return longstepT, latstepT
}

func (service *ServiceT) generate() {
	longstep, latstep := getSteps(service.PointA, service.LongitudeStep, service.LatitudeStep)
	//longstep_s, latstep_s := FloatSafe{fmt.Sprintf("%v", longstep)}, FloatSafe{fmt.Sprintf("%v", latstep)}
	//fromLatitude, toLatitude := FloatSafe{fmt.Sprintf("%v", service.PointA.Latitude)}, FloatSafe{fmt.Sprintf("%v", service.PointB.Latitude)}
	//fromLongitude, toLongitude := FloatSafe{fmt.Sprintf("%v", service.PointA.Longitude)}, FloatSafe{fmt.Sprintf("%v", service.PointB.Longitude)}
	fromLatitude, fromLongitude := service.PointA.Latitude, service.PointA.Longitude
	toLatitude, toLongitude := service.PointB.Latitude, service.PointB.Longitude

	//if fromLatitude.getValue() > toLatitude.getValue() {
	if fromLatitude > toLatitude {
		//for fromLatitude.getValue() > toLatitude.getValue() {
		for fromLatitude > toLatitude {
			tmp := make([]*Point, 0)
			//if fromLongitude.getValue() > toLongitude.getValue() {
			if fromLongitude > toLongitude {
				//for fromLongitude.getValue() > toLongitude.getValue() {
				for fromLongitude > toLongitude {
					//tmp = append(tmp, newPoint(fromLatitude.getValue(), fromLongitude.getValue(), -1))
					tmp = append(tmp, newPoint(fromLatitude, fromLongitude, -1))
					fromLongitude -= longstep
					//fromLongitude.diff(&longstep_s)
				}
			} else {
				//for fromLongitude.getValue() < toLongitude.getValue() {
				for fromLongitude < toLongitude {
					tmp = append(tmp, newPoint(fromLatitude, fromLongitude, -1))
					fromLongitude += longstep
					//fromLongitude.add(&longstep_s)
				}
			}
			//fromLongitude.setValue(service.PointA.Longitude)
			fromLongitude = service.PointA.Longitude
			fromLatitude -= latstep
			//fromLatitude.diff(&latstep_s)
			service.Area = append(service.Area, tmp)
		}
	} else {
		//for fromLatitude.getValue() < toLatitude.getValue() {
		for fromLatitude < toLatitude {
			tmp := make([]*Point, 0)
			//if fromLongitude.getValue() > toLongitude.getValue() {
			if fromLongitude > toLongitude {
				for fromLongitude > toLongitude {
					//tmp = append(tmp, newPoint(fromLatitude.getValue(), fromLongitude.getValue(), -1))
					tmp = append(tmp, newPoint(fromLatitude, fromLongitude, -1))
					fromLongitude -= longstep
					//fromLongitude.diff(&longstep_s)
				}
			} else {
				//for fromLongitude.getValue() < toLongitude.getValue() {
				for fromLongitude < toLongitude {
					//tmp = append(tmp, newPoint(fromLatitude.getValue(), fromLongitude.getValue(), -1))
					tmp = append(tmp, newPoint(fromLatitude, fromLongitude, -1))
					fromLongitude += longstep
					//fromLongitude.add(&longstep_s)
				}
			}
			//fromLongitude.setValue(service.PointA.Longitude)
			fromLongitude = service.PointA.Longitude
			fromLatitude += latstep
			//fromLatitude.add(&latstep_s)
			service.Area = append(service.Area, tmp)
		}
	}
}
