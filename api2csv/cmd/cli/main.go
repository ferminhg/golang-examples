package main

import (
	"fmt"

	"github.com/CodelyTV/golang-examples/api2csv/internal/cli"
	"github.com/CodelyTV/golang-examples/api2csv/internal/fetching"
	"github.com/CodelyTV/golang-examples/api2csv/internal/storaging"

	// inmemory "github.com/CodelyTV/golang-examples/api2csv/internal/storage/inmem"
	csv "github.com/CodelyTV/golang-examples/api2csv/internal/storage/csv"
	swapi "github.com/CodelyTV/golang-examples/api2csv/internal/storage/swapi"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("wop wop wop 🥤")

	repo := swapi.NewRepository()
	fileRepo := csv.NewStorageRepo()

	fetchingService := fetching.NewService(repo)
	storagingService := storaging.NewService(fileRepo)
	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cli.InitETLCmd(fetchingService, storagingService))
	rootCmd.Execute()
}
