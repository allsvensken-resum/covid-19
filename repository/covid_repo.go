package repository

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

type Response struct {
	Data []CovidPatient
}

type covidRepo struct {
	url string
}

func NewCovidRepository(url string) ICovidRepo {
	return covidRepo{url: url}
}

func (t covidRepo) GetAllCovidPatient() ([]CovidPatient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(t.url)
	if err != nil {
		return nil, err
	}

	var responseModel Response
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseModel); err != nil {
		return nil, err
	}

	return responseModel.Data, nil
}
