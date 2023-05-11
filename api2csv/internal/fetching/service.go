package fetching

import (
	etlcli "github.com/CodelyTV/golang-examples/api2csv/internal"
)

type Service interface {
	FetchPeople() ([]etlcli.People, error)
}

type service struct {
	pR etlcli.PeopleRepo
}

func NewService(r etlcli.PeopleRepo) Service {
	return &service{r}
}

func (s *service) FetchPeople() ([]etlcli.People, error) {
	return s.pR.GetPeople()
}
