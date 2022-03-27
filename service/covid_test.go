package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecideAgeRange(t *testing.T) {
	type testCase struct {
		name     string
		age      *int
		expected string
	}

	ages := []int{19, 31, 65}

	cases := []testCase{
		{
			name:     "should be in range 0 - 30",
			age:      &ages[0],
			expected: "0-30",
		},
		{
			name:     "should be in range 31 - 60",
			age:      &ages[1],
			expected: "31-60",
		},
		{
			name:     "should be in range 61+",
			age:      &ages[2],
			expected: "61+",
		},
		{
			name:     "should be in range N/A",
			age:      nil,
			expected: "N/A",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ageRange := decideAgeRange(c.age)
			assert.Equal(t, c.expected, ageRange)
		})
	}
}

func TestGroupPatientsByProvince(t *testing.T) {
	type testCase struct {
		name     string
		province *string
		expected int
	}

	patientsGroupByProvince := make(map[string]int)
	patientsGroupByProvince["Chanthaburi"] = 10
	provinces := []string{"Chanthaburi", "Buriram"}

	cases := []testCase{
		{name: "should increase by 1", province: &provinces[0], expected: 11},
		{name: "should init at 1", province: &provinces[1], expected: 1},
	}

	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			groupPatientsByProvince(c.province, patientsGroupByProvince)
			province := provinces[i]
			actual := patientsGroupByProvince[province]
			assert.Equal(t, c.expected, actual)
		})
	}

	expected := len(patientsGroupByProvince)

	t.Run("should not increase number of provinces", func(t *testing.T) {
		groupPatientsByProvince(nil, patientsGroupByProvince)
		actual := len(patientsGroupByProvince)
		assert.Equal(t, expected, actual)
	})

}

func TestGroupPatientsByAge(t *testing.T) {
	type testCase struct {
		name     string
		ageRange string
		expected int
	}

	patientsGroupByAge := make(map[string]int)
	patientsGroupByAge["0-30"] = 20

	cases := []testCase{
		{name: "should increment at range 0-30 by 1", ageRange: "0-30", expected: 21},
		{name: "should init at range 31-60 at 1", ageRange: "31-60", expected: 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			groupPatientsByAge(c.ageRange, patientsGroupByAge)
			ageRange := c.ageRange
			actual := patientsGroupByAge[ageRange]
			assert.Equal(t, c.expected, actual)
		})
	}
}
