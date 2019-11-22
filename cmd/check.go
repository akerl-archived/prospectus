package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/akerl/prospectus/plugin"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for state changes",
	RunE:  checkRunner,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	f := checkCmd.Flags()
	f.BoolP("all", "a", false, "Show all items, regardless of state")
	f.Bool("json", false, "Print output as JSON")
}

func checkRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	flagAll, err := flags.GetBool("all")
	if err != nil {
		return err
	}
	flagJSON, err := flags.GetBool("json")
	if err != nil {
		return err
	}

	params := []string{"."}
	if len(args) != 0 {
		params = args
	}

	as, err := plugin.NewSet(params)
	if err != nil {
		return err
	}
	results := as.Check()
	if err != nil {
		return err
	}
	if !flagAll {
		newResultSet := plugin.ResultSet{}
		for _, item := range results {
			if !item.Matches {
				newResultSet = append(newResultSet, item)
			}
		}
		results = newResultSet
	}

	var output string
	if flagJSON {
		outputBytes, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return err
		}
		output = string(outputBytes)
	} else {
		output = results.String()
	}
	fmt.Println(output)
	return nil
}
