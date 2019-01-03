package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/akerl/prospectus-ng/prospectus"

	"github.com/spf13/cobra"
)

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

	if len(args) == 0 {
		args = []string{"."}
	}

	checks := prospectus.Checks{}
	err = checks.Populate(args)
	if err != nil {
		return err
	}
	results, err := checks.Execute()
	if err != nil {
		return err
	}

	var show prospectus.Results
	if flagAll {
		show = results
	} else {
		for _, i := range results {
			if i.Matches() {
				show = append(show, i)
			}
		}
	}

	var output string
	if flagJSON {
		data, err := json.MarshalIndent(show, "", "  ")
		if err != nil {
			return err
		}
		output = string(data)
	} else {
		var b strings.Builder
		for _, i := range show {
			_, err := b.WriteString(i.String())
			if err != nil {
				return err
			}
			_, err = b.WriteString("\n")
			if err != nil {
				return err
			}
		}
		output = b.String()
	}
	fmt.Println(output)
	return nil
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
