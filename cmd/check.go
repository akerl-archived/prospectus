package cmd

import (
	"fmt"

	"github.com/akerl/prospectus/v3/runner"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for state changes",
	RunE:  checkRunner,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolP("all", "a", false, "Show all items, regardless of state")
}

func checkRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	flagAll, err := flags.GetBool("all")
	if err != nil {
		return err
	}

	dirs := []string{"."}
	if len(args) != 0 {
		dirs = args
	}

	list, err := runner.NewSet(dirs)
	if err != nil {
		return err
	}

	results := as.Check()
	if !flagAll {
		results = changedResults(results)
	}

	fmt.Println(results.String())
	return nil
}

func changedResults(rs runner.Set) runner.Set {
	newResults := runner.Set{}
	for _, item := range rs {
		if !item.Matches {
			newResults = append(newResults, item)
		}
	}
	return newResults
}
