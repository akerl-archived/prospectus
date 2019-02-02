package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "prospectus",
	Short:         "Tool and DSL for checking expected vs actual state",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute function is the entrypoint for the CLI
func Execute() error {
	return rootCmd.Execute()
}
