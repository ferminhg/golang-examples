package storaging

import (
	etlcli "github.com/CodelyTV/golang-examples/api2csv/internal"
)

type Service interface {
	StorageService(people []etlcli.People) error
}

type service struct {
	sR etlcli.StorageRepo
}

func NewService(r etlcli.StorageRepo) Service {
	return &service{r}
}

func (s *service) StorageService(people []etlcli.People) error {
	s.sR.StoragePeople(people)
	return nil
}
