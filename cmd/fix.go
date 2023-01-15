package cmd

import (
	"fmt"

	"github.com/akerl/prospectus/v3/runner"

	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Attempt to fix items where expected state differs from actual state",
	RunE:  fixRunner,
}

func init() {
	rootCmd.AddCommand(fixCmd)
}

func fixRunner(_ *cobra.Command, args []string) error {
	dirs := []string{"."}
	if len(args) != 0 {
		dirs = args
	}

	list, err := runner.NewSet(dirs)
	if err != nil {
		return err
	}

	results := as.Check().Fix()

	fmt.Println(results.String())
	return nil
}
