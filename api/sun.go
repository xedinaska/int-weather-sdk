package api

import "time"

type SunInfoRequest struct {
	Date      string  `json:"date"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SunInfoResponse struct {
	Rise time.Time `json:"rise"`
	Set  time.Time `json:"set"`
}
