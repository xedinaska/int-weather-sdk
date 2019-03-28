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

type DateWeatherHandler struct {
	serviceName string
	logger      *logrus.Entry
	integration sdk.Integration
}

func NewDateWeatherHandler(serviceName string, logger *logrus.Entry, integration sdk.Integration) *DateWeatherHandler {
	return &DateWeatherHandler{
		serviceName: serviceName,
		logger:      logger,
		integration: integration,
	}
}

func (h *DateWeatherHandler) GetWeatherForDate(req *restful.Request, rsp *restful.Response) {

	ctx := req.Attribute("ctx").(context.Context)
	logger := logrus.Entry{}

	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		logger.Errorf("failed to unmarshal request body %v", err.Error())
		return
	}
	defer req.Request.Body.Close()

	dateRequest := &api.DateWeatherRequest{}
	if err := json.Unmarshal(body, dateRequest); err != nil {
		logger.WithField("request", fmt.Sprintf("%+v", req)).Errorf("[%s] invalid incoming DateWeather request", h.serviceName)
		rsp.WriteHeaderAndEntity(422, &api.Error{
			Status:  422,
			Message: err.Error(),
		})
		return
	}

	response, err := h.integration.GetWeatherForDate(ctx, dateRequest)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"request":  fmt.Sprintf("%+v", dateRequest),
			"response": fmt.Sprintf("%+v", response),
		}).Errorf("[%s] DateWeather error: %s", h.serviceName, err.Error())

		rsp.WriteHeaderAndEntity(http.StatusInternalServerError, &api.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	rsp.WriteHeaderAndEntity(http.StatusOK, response)
}
