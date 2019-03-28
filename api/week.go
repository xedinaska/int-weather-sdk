package api

type WeekWeatherRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WeekWeatherResponse struct {
	Weather   []*DateWeatherResponse `json:"weather"`
	WeekStart string                 `json:"week_start"`
	WeekEnd   string                 `json:"week_end"`
}
