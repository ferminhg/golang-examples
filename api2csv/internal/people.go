package etlcli

// https://swapi.dev/
type People struct {
	Name   string
	Gender string
	Height int
	Mass   int
}

type PeopleRepo interface {
	GetPeople() ([]People, error)
}

func NewPeople(Name string, Gender string, Height int, Mass int) People {
	return People{
		Name:   Name,
		Gender: Gender,
		Height: Height,
		Mass:   Mass,
	}
}
