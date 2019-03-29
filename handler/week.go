package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk"
	"github.com/xedinaska/int-weather-sdk/api"
)

type WeekWeatherHandler struct {
	serviceName string
	logger      *logrus.Entry
	integration sdk.Integration
}

func NewWeekWeatherHandler(serviceName string, logger *logrus.Entry, integration sdk.Integration) *WeekWeatherHandler {
	return &WeekWeatherHandler{
		serviceName: serviceName,
		logger:      logger,
		integration: integration,
	}
}

func (h *WeekWeatherHandler) GetWeekWeather(req *restful.Request, rsp *restful.Response) {

	ctx := req.Attribute("ctx").(context.Context)

	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		h.logger.Errorf("failed to unmarshal request body %v", err.Error())
		return
	}
	defer req.Request.Body.Close()

	weekRequest := &api.WeekWeatherRequest{}
	if err := json.Unmarshal(body, weekRequest); err != nil {
		h.logger.WithField("request", fmt.Sprintf("%+v", req)).Errorf("[%s] invalid incoming WeekWeather request", h.serviceName)
		rsp.WriteHeaderAndEntity(422, &api.Error{
			Status:  422,
			Message: err.Error(),
		})
		return
	}

	response, err := h.integration.GetWeekWeather(ctx, weekRequest)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"request":  fmt.Sprintf("%+v", weekRequest),
			"response": fmt.Sprintf("%+v", response),
		}).Errorf("[%s] WeekWeather error: %s", h.serviceName, err.Error())

		rsp.WriteHeaderAndEntity(http.StatusInternalServerError, &api.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	rsp.WriteHeaderAndEntity(http.StatusOK, response)
}
