package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var help = &cobra.Command{
	Use:   "help",
	Short: "Print the help for Pilot",
	Long:  "All software has help. This is Pilot's",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Pull from somewhere
		fmt.Println("Pilot's version is currently alpha")
	},
}

func init() {
	root.AddCommand(help)
}
