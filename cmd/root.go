package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "pilot",
	Short: "Pilot is the command line tools that helps you to maintain and setup your project seamlessly.",
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("pilot command execution failed because of an error - '%s'", err)
	}
}
