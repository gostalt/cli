package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gostalt",
	Short: "Run commands for the Gostalt framework.",
}

func init() {
	root.AddCommand(new)
}

// Execute is the entry point to the CLI app.
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
