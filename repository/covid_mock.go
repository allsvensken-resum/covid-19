package repository

import "github.com/stretchr/testify/mock"

type covidRepoMock struct {
	mock.Mock
}

func NewCovidRepositoryMock() *covidRepoMock {
	return &covidRepoMock{}
}

func (c *covidRepoMock) GetAllCovidPatient() ([]CovidPatient, error) {
	args := c.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]CovidPatient), args.Error(1)
}
