package api

type DateWeatherRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Date      string  `json:"date"`
}

type DateWeatherResponse struct {
	StateName  string  `json:"state_name"`
	MinTemp    float64 `json:"min_temp"`
	MaxTemp    float64 `json:"max_temp"`
	WindSpeed  float64 `json:"wind_speed"`
	Humidity   int     `json:"humidity"`
	Visibility float64 `json:"visibility"`
	Date       string  `json:"date"`
}
