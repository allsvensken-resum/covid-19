package service

import (
	"log"

	"github.com/suppakorn-dev/lmwn-covid-19/errs"
	"github.com/suppakorn-dev/lmwn-covid-19/repository"
)

type covidSrv struct {
	covidRepo repository.ICovidRepo
}

func NewCovidService(repo repository.ICovidRepo) ICovidSrv {
	return covidSrv{covidRepo: repo}
}

func (t covidSrv) GetCovidPatientSummary() (PatientSummaryResp, error) {
	patients, err := t.covidRepo.GetAllCovidPatient()
	if err != nil {
		log.Println(err)
		return PatientSummaryResp{}, errs.NewNotFoundError()
	}

	patientsGroupByProvince := make(map[string]int)
	patientsGroupByAge := make(map[string]int)

	for _, patient := range patients {
		groupPatientsByProvince(patient.ProvinceEn, patientsGroupByProvince)
		ageRange := decideAgeRange(patient.Age)
		groupPatientsByAge(ageRange, patientsGroupByAge)
	}

	return PatientSummaryResp{Province: patientsGroupByProvince, AgeGroup: patientsGroupByAge}, nil
}

func decideAgeRange(age *int) string {
	if age == nil {
		return "N/A"
	} else if *age >= 0 && *age <= 30 {
		return "0-30"
	} else if *age >= 31 && *age <= 60 {
		return "31-60"
	}
	return "61+"
}

func groupPatientsByAge(ageRange string, patientsGroupByAge map[string]int) {
	if count, ok := patientsGroupByAge[ageRange]; ok {
		patientsGroupByAge[ageRange] = count + 1
		return
	}
	patientsGroupByAge[ageRange] = 1
}

func groupPatientsByProvince(province *string, patientsGroupByProvince map[string]int) {
	if province == nil {
		return
	}

	if count, ok := patientsGroupByProvince[*province]; ok {
		patientsGroupByProvince[*province] = count + 1
		return
	}
	patientsGroupByProvince[*province] = 1
}
