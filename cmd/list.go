package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/akerl/prospectus/v2/plugin"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list checks",
	RunE:  listRunner,
}

func init() {
	rootCmd.AddCommand(listCmd)
	f := listCmd.Flags()
	f.Bool("json", false, "Print output as JSON")
}

func listRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
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

	var output string
	if flagJSON {
		outputBytes, err := json.MarshalIndent(as, "", "  ")
		if err != nil {
			return err
		}
		output = string(outputBytes)
	} else {
		output = as.String()
	}
	fmt.Println(output)
	return nil
}
