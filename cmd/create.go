package cmd

import (
	"log"

	"github.com/devterre/pilot/pkg/create"
	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  `Pilot will have you structure the code, that is easier to maintain and scale.`,
	Run: func(cmd *cobra.Command, args []string) {
		createCommandFunc()
	},
}

func init() {
	root.AddCommand(createCommand)
}

func createCommandFunc() {
	err := create.CreateNewFolder()

	if err != nil {
		log.Fatalf("Exiting the process during initial stage.", err)
	}
}
