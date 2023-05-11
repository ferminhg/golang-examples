package cli

import (
	"fmt"

	"github.com/CodelyTV/golang-examples/api2csv/internal/fetching"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func InitETLCmd(service fetching.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "people",
		Short: "Print people list",
		Run:   runETLFn(service),
	}
	return cmd
}

func runETLFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		people, _ := service.FetchPeople()
		fmt.Println(people)
	}
}
