package inmemory

import (
	peoplecli "github.com/CodelyTV/golang-examples/api2csv/internal"
)

type repository struct {
}

func NewRepository() peoplecli.PeopleRepo {
	return &repository{}
}

func (r *repository) GetPeople() ([]peoplecli.People, error) {
	people := []peoplecli.People{}
	people = append(people,
		peoplecli.NewPeople("John Doe", "X"),
		peoplecli.NewPeople("Beka", "Female"),
		peoplecli.NewPeople("Maya", "Female"),
	)
	return people, nil
}
