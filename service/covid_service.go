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

func (t covidSrv) GetCovidPatientSummary() (PatientSummaryResp, error) {
	patients, err := t.covidRepo.GetAllCovidPatient()
	if err != nil {
		log.Fatal(err)
		return PatientSummaryResp{}, err
	}

	patientsGroupByProvince := make(map[string]int)
	ageRanges := []string{"0-30", "31-60", "61+", "N/A"}
	patientsGroupByAge := initAgeRangesMap(ageRanges)

	for _, patient := range patients {

		if patient.Age == nil {
			patientsGroupByAge["N/A"] = patientsGroupByAge["N/A"] + 1
		} else if *patient.Age >= 0 && *patient.Age <= 30 {
			patientsGroupByAge["0-30"] = patientsGroupByAge["0-30"] + 1
		} else if *patient.Age >= 31 && *patient.Age <= 60 {
			patientsGroupByAge["31-60"] = patientsGroupByAge["31-60"] + 1
		} else if *patient.Age >= 61 {
			patientsGroupByAge["61+"] = patientsGroupByAge["61+"] + 1

		}
		if patient.ProvinceEn == nil {
			continue
		}

		if count, ok := patientsGroupByProvince[*patient.ProvinceEn]; ok {
			patientsGroupByProvince[*patient.ProvinceEn] = count + 1
		} else {
			patientsGroupByProvince[*patient.ProvinceEn] = 1
		}
	}

	return PatientSummaryResp{Province: patientsGroupByProvince, AgeGroup: patientsGroupByAge}, nil
}

func initAgeRangesMap(ageRanges []string) map[string]int {
	patientsGroupByAges := make(map[string]int)

	for _, ageRange := range ageRanges {
		patientsGroupByAges[ageRange] = 0
	}

	return patientsGroupByAges
}
