package service

import (
	"log"

	"github.com/suppakorn-dev/lmwn-covid-19/repository"
)

type covidSrv struct {
	covidRepo repository.ICovidRepo
}

func NewCovidService(repo repository.ICovidRepo) ICovidSrv {
	return covidSrv{covidRepo: repo}
}

func (t covidSrv) GetCovidPatientSummary() {
	patients, err := t.covidRepo.GetAllCovidPatient()
	if err != nil {
		log.Fatal(err)
	}
	_ = patients
}
