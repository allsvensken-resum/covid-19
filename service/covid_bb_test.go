package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suppakorn-dev/lmwn-covid-19/errs"
	"github.com/suppakorn-dev/lmwn-covid-19/repository"
	"github.com/suppakorn-dev/lmwn-covid-19/service"
)

func TestGetCovidPatientSummary(t *testing.T) {
	t.Run("should immediately return not found error", func(t *testing.T) {
		repo := repository.NewCovidRepositoryMock()
		repo.On("GetAllCovidPatient").Return(nil, errors.New("Error"))
		service := service.NewCovidService(repo)

		_, err := service.GetCovidPatientSummary()

		assert.ErrorIs(t, err, errs.NewNotFoundError())
	})

	t.Run("should return two grouping result", func(t *testing.T) {
		repo := repository.NewCovidRepositoryMock()
		age := 20
		provinceEn := "Chanthaburi"

		repo.On("GetAllCovidPatient").Return([]repository.CovidPatient{
			{Age: &age, ProvinceEn: &provinceEn},
		}, nil)

		service := service.NewCovidService(repo)

		actual, _ := service.GetCovidPatientSummary()
		assert.Equal(t, 1, actual.Province[provinceEn])
		assert.Equal(t, 1, actual.AgeGroup["0-30"])
	})
}
