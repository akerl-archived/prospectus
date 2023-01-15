package cmd

import (
	"fmt"

	"github.com/akerl/prospectus/v3/runner"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list checks",
	RunE:  listRunner,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listRunner(_ *cobra.Command, args []string) error {
	dirs := []string{"."}
	if len(args) != 0 {
		dirs = args
	}

	list, err := runner.NewSet(dirs)
	if err != nil {
		return err
	}

	fmt.Println(list.String())
	return nil
}
