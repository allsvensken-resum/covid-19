package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suppakorn-dev/lmwn-covid-19/errs"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

type CovidHandler struct {
	covidService service.ICovidSrv
}

func NewCovidHandler(service service.ICovidSrv) CovidHandler {
	return CovidHandler{covidService: service}
}

func (t CovidHandler) GetCovidPatientSummary(c *gin.Context) {
	summary, err := t.covidService.GetCovidPatientSummary()

	if err, ok := err.(errs.AppError); ok {
		c.JSON(err.Code, gin.H{"Message": err.Message})
		return
	}

	if err != nil {
		unexpectedErr, _ := errs.NewUnexpectedError().(errs.AppError)
		c.JSON(unexpectedErr.Code, gin.H{"Message": unexpectedErr.Message})
		return
	}

	c.JSON(http.StatusOK, summary)
}
