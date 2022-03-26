package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suppakorn-dev/lmwn-covid-19/handler"
	"github.com/suppakorn-dev/lmwn-covid-19/repository"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

func main() {
	router := gin.Default()
	url := "https://static.wongnai.com/devinterview/covid-cases.json"
	covidRepo := repository.NewCovidRepository(url)
	covidSrv := service.NewCovidService(covidRepo)
	covidSrv.GetCovidPatientSummary()
	covidHandler := handler.NewCovidHandler(covidSrv)

	router.GET("/covid/summary", covidHandler.GetCovidPatientSummary)
	router.Run(":80")
}
