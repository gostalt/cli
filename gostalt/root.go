package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gostalt",
	Short: "Gostalt is a really nice Go Framework",
}

func init() {
	root.AddCommand(new)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
