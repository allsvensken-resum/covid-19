package service

import "github.com/stretchr/testify/mock"

type CovidSrvMock struct {
	mock.Mock
}

func NewCovidSrvMock() *CovidSrvMock {
	return &CovidSrvMock{}
}

func (c *CovidSrvMock) GetCovidPatientSummary() (PatientSummaryResp, error) {
	args := c.Called()
	return args.Get(0).(PatientSummaryResp), args.Error(1)
}
