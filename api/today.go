package api

type TodayWeatherRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TodayWeatherResponse struct {
	StateName  string  `json:"state_name"`
	MinTemp    float64 `json:"min_temp"`
	MaxTemp    float64 `json:"max_temp"`
	WindSpeed  float64 `json:"wind_speed"`
	Humidity   int     `json:"humidity"`
	Visibility float64 `json:"visibility"`
}
