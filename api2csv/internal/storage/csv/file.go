package csv

import (
	"encoding/csv"
	"log"
	"os"

	peoplecli "github.com/CodelyTV/golang-examples/api2csv/internal"
)

const (
	tmpCSVFile = "/tmp/people.csv"
)

type storage struct {
	csvFile string
}

func NewStorageRepo() peoplecli.StorageRepo {
	return &storage{csvFile: tmpCSVFile}
}

func (r *storage) StoragePeople(people []peoplecli.People) error {
	file, err := os.Create(r.csvFile)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(file)

	_ = csvWriter.Write([]string{"Name", "Gender"})
	for _, person := range people {
		_ = csvWriter.Write([]string{person.Name, person.Gender})
	}

	csvWriter.Flush()
	file.Close()
	return nil
}
