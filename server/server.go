package server

import (
	"fmt"
	
	microweb "github.com/micro/go-web"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk"
	"github.com/xedinaska/int-weather-sdk/config"
	"github.com/xedinaska/int-weather-sdk/filter"
	"github.com/xedinaska/int-weather-sdk/handler"
)

const port = 8091

type Server struct {
	WebRouter      *restful.Container
	WebService     microweb.Service
	RestfulService *restful.WebService
}

func Create(serviceName, serviceVersion string, integration sdk.Integration, conf *config.Config) *Server {
	// Create a restful web-service:
	restfulService := new(restful.WebService).
		Path("/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// Create a restful web-container (this will be our router):
	webRouter := restful.NewContainer()

	// Create a new Micro WEB service:
	webService := microweb.NewService(
		microweb.Name(serviceName),
		microweb.Version(serviceVersion),
		microweb.Address(fmt.Sprintf(":%d", port)),
		microweb.Metadata(map[string]string{
			"type": "integration",
		}),

		// Handle with our restful router:
		microweb.Handler(webRouter),

		// Add runtime flags:
		microweb.Flags(),
	)

	// Init:
	webService.Init()

	logger := logrus.New().WithFields(
		logrus.Fields{
			"logger":      "main",
			"integration": serviceName,
			"version":     serviceVersion,
		})

	// Create context builder filter
	contextBuilder, err := filter.NewContextBuilder(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Add the filters
	restfulService.Filter(contextBuilder.CreateContext)

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
		WebService:     webService,
		RestfulService: restfulService,
	}
}
