package etlcli

type People struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type PeopleRepo interface {
	GetPeople() ([]People, error)
}

type StorageRepo interface {
	StoragePeople(people []People) error
}

func NewPeople(Name string, Gender string) People {
	return People{
		Name:   Name,
		Gender: Gender,
	}
}
