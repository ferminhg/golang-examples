package swapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	peoplecli "github.com/CodelyTV/golang-examples/api2csv/internal"
)

const (
	swapiPeopleUrl = "https://swapi.dev/api/people"
)

type repository struct {
	url string
}

func NewRepository() peoplecli.PeopleRepo {
	return &repository{url: swapiPeopleUrl}
}

type swapiPeopleUrlResponse struct {
	Results []peoplecli.People `json:"results"`
}

func (r *repository) GetPeople() ([]peoplecli.People, error) {
	response, err := http.Get(fmt.Sprintf("%v", r.url))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp swapiPeopleUrlResponse
	err = json.Unmarshal(contents, &resp)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return resp.Results, nil
}
