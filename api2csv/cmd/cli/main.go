package main

import (
	"fmt"

	"github.com/CodelyTV/golang-examples/api2csv/internal/cli"
	"github.com/CodelyTV/golang-examples/api2csv/internal/fetching"
	inmemory "github.com/CodelyTV/golang-examples/api2csv/internal/storage"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("wopwop")

	// var repo peoplecli.PeopleRepo
	// repo = inmemory.NewRepository()
	repo := inmemory.NewRepository()

	fetchingService := fetching.NewService(repo)
	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cli.InitETLCmd(fetchingService))
	rootCmd.Execute()
}
