package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/suppakorn-dev/lmwn-covid-19/errs"
	"github.com/suppakorn-dev/lmwn-covid-19/handler"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

func TestGetCovidPatientSummary(t *testing.T) {
	rPath := "/covid/summary"

	t.Run("should return not found error when getting app error", func(t *testing.T) {
		covidSrv := service.NewCovidSrvMock()
		covidSrv.On("GetCovidPatientSummary").Return(service.PatientSummaryResp{}, errs.NewNotFoundError())
		covidHandler := handler.NewCovidHandler(covidSrv)
		router := gin.Default()
		router.GET(rPath, covidHandler.GetCovidPatientSummary)
		req, _ := http.NewRequest("GET", rPath, nil)
		w := httptest.NewRecorder()
		expected := http.StatusNotFound

		router.ServeHTTP(w, req)

		assert.Equal(t, expected, w.Code)
	})

	t.Run("should return not internal server error when getting unexpected error", func(t *testing.T) {
		covidSrv := service.NewCovidSrvMock()
		covidSrv.On("GetCovidPatientSummary").Return(service.PatientSummaryResp{}, errors.New("Unexpected"))
		covidHandler := handler.NewCovidHandler(covidSrv)
		router := gin.Default()
		router.GET(rPath, covidHandler.GetCovidPatientSummary)
		req, _ := http.NewRequest("GET", rPath, nil)
		w := httptest.NewRecorder()
		expected := http.StatusInternalServerError

		router.ServeHTTP(w, req)

		assert.Equal(t, expected, w.Code)
	})

	t.Run("should return status OK if not getting error", func(t *testing.T) {
		covidSrv := service.NewCovidSrvMock()
		covidSrv.On("GetCovidPatientSummary").Return(service.PatientSummaryResp{}, nil)
		covidHandler := handler.NewCovidHandler(covidSrv)
		router := gin.Default()
		router.GET(rPath, covidHandler.GetCovidPatientSummary)
		req, _ := http.NewRequest("GET", rPath, nil)
		w := httptest.NewRecorder()
		expected := http.StatusOK

		router.ServeHTTP(w, req)

		assert.Equal(t, expected, w.Code)
	})

}
