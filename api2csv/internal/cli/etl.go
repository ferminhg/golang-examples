package cli

import (
	"github.com/CodelyTV/golang-examples/api2csv/internal/fetching"
	"github.com/CodelyTV/golang-examples/api2csv/internal/storaging"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func InitETLCmd(fetchingService fetching.Service, storageService storaging.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "people",
		Short: "Print people list",
		Run:   runETLFn(fetchingService, storageService),
	}
	return cmd
}

func runETLFn(fetchingService fetching.Service, storageService storaging.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		people, _ := fetchingService.FetchPeople()
		storageService.StorageService(people)
	}
}
