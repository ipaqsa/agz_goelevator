package server

type response struct {
	Comment string `json:"comment"`
}

type requestT struct {
	A             string `json:"A"`
	B             string `json:"B"`
	LongitudeStep string `json:"Longitude-step"`
	LatitudeStep  string `json:"Latitude-step"`
	Source        string `json:"source"`
}
