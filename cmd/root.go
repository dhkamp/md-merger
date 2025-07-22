package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "md-merger",
		Short: "md-merger is a CLI tool to merge markdown files",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to md-merger")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
