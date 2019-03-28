package server

import (
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk"
	"github.com/xedinaska/int-weather-sdk/handler"
)

type Server struct {
	WebRouter      *restful.Container
	RestfulService *restful.WebService
}

func Create(serviceName, serviceVersion string, integration sdk.Integration) *Server {
	// Create a restful web-service:
	restfulService := new(restful.WebService).
		Path("/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// Create a restful web-container (this will be our router):
	webRouter := restful.NewContainer()

	logger := logrus.New().WithFields(
		logrus.Fields{
			"logger":      "main",
			"integration": serviceName,
			"version":     serviceVersion,
		})

	todayWeatherHandler := handler.NewTodayWeatherHandler(serviceName, logger, integration)
	restfulService.Route(restfulService.POST("/today").To(todayWeatherHandler.GetTodayWeather))

	weekWeatherHandler := handler.NewWeekWeatherHandler(serviceName, logger, integration)
	restfulService.Route(restfulService.POST("/week").To(weekWeatherHandler.GetWeekWeather))

	dateWeatherHandler := handler.NewDateWeatherHandler(serviceName, logger, integration)
	restfulService.Route(restfulService.POST("/date").To(dateWeatherHandler.GetWeatherForDate))

	sunInfoHandler := handler.NewSunInfoHandler(serviceName, logger, integration)
	restfulService.Route(restfulService.POST("/sun").To(sunInfoHandler.GetSunInfo))

	webRouter.Add(restfulService)

	return &Server{
		WebRouter:      webRouter,
		RestfulService: restfulService,
	}
}
