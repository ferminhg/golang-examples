package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

func runFn(elements map[string]string) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(elements[id])
		} else {
			fmt.Println(elements)
		}
	}
}
