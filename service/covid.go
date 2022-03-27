package service

type PatientSummaryResp struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

type ICovidSrv interface {
	GetCovidPatientSummary() (PatientSummaryResp, error)
}
