package cmd

import (
	"fmt"

	"github.com/akerl/prospectus-ng/prospectus"

	"github.com/spf13/cobra"
)

func checkRunner(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		args = []string{"."}
	}
	
	results = prospectus.Results{}
	for _, x := args {
		// TODO: make this run in parallel
	}
}

var checkCmd = &cobra.Command{
	Use:   "check ROLENAME",
	Short: "Check for state changes",
	RunE:  checkRunner,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	f := checkCmd.Flags()
	f.BoolP("all", "a", false, "Show all items, regardless of state")
	f.Bool("json", false, "Print output as JSON")
}
