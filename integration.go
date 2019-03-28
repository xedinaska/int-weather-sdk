package sdk

import (
	"context"
	"github.com/xedinaska/int-weather-sdk/api"
)

// Integration is the common interface all DMS integrations must implement
type Integration interface {
	GetTodayWeather(ctx context.Context, req *api.TodayWeatherRequest) (*api.TodayWeatherResponse, error)
	GetWeekWeather(ctx context.Context, req *api.WeekWeatherRequest) (*api.WeekWeatherResponse, error)
	GetWeatherForDate(ctx context.Context, req *api.DateWeatherRequest) (*api.DateWeatherResponse, error)
	GetSunInfo(ctx context.Context, req *api.SunInfoRequest) (*api.SunInfoResponse, error)
}
