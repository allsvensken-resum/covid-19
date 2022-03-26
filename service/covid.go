package service

type PatientSummaryResp struct {
	Province map[string]int
	AgeGroup map[string]int
}

type ICovidSrv interface {
	GetCovidPatientSummary()
}
