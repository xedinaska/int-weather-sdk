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

type TodayWeatherHandler struct {
	serviceName string
	logger      *logrus.Entry
	integration sdk.Integration
}

func NewTodayWeatherHandler(serviceName string, logger *logrus.Entry, integration sdk.Integration) *TodayWeatherHandler {
	return &TodayWeatherHandler{
		serviceName: serviceName,
		logger:      logger,
		integration: integration,
	}
}

func (h *TodayWeatherHandler) GetTodayWeather(req *restful.Request, rsp *restful.Response) {

	ctx := req.Attribute("ctx").(context.Context)

	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		h.logger.Errorf("failed to unmarshal request body %v", err.Error())
		return
	}
	defer req.Request.Body.Close()

	todayRequest := &api.TodayWeatherRequest{}
	if err := json.Unmarshal(body, todayRequest); err != nil {
		h.logger.WithField("request", fmt.Sprintf("%+v", req)).Errorf("[%s] invalid incoming TodayWeather request", h.serviceName)
		rsp.WriteHeaderAndEntity(422, &api.Error{
			Status:  422,
			Message: err.Error(),
		})
		return
	}

	response, err := h.integration.GetTodayWeather(ctx, todayRequest)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"request":  fmt.Sprintf("%+v", todayRequest),
			"response": fmt.Sprintf("%+v", response),
		}).Errorf("[%s] TodayWeather error: %s", h.serviceName, err.Error())

		rsp.WriteHeaderAndEntity(http.StatusInternalServerError, &api.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	rsp.WriteHeaderAndEntity(http.StatusOK, response)
}
