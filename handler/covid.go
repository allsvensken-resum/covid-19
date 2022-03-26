package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

type CovidHandler struct {
	covidService service.ICovidSrv
}

func NewCovidHandler(service service.ICovidSrv) CovidHandler {
	return CovidHandler{covidService: service}
}

func (t CovidHandler) GetCovidPatientSummary(c *gin.Context) {

}
