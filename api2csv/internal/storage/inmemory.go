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
		peoplecli.NewPeople("John Doe", "X", 0, 0),
		peoplecli.NewPeople("Beka", "Female", 10, 8),
		peoplecli.NewPeople("Maya", "Female", 5, 6),
	)
	return people, nil
}
