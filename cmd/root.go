/*
Copyright © 2026 KYLE LINDSAY

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "issues",
	Short: "A CLI tool for managing project issues",
	Long: `A CLI tool that allows the creation, tracking, modification and deletion of issues directly in the terminal`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.CompletionOptions.DisableDefaultCmd = true
}


