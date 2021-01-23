package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/akerl/prospectus/v2/plugin"

	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Attempt to fix items where expected state differs from actual state",
	RunE:  fixRunner,
}

func init() {
	rootCmd.AddCommand(fixCmd)
	f := fixCmd.Flags()
	f.Bool("json", false, "Print output as JSON")
}

func fixRunner(cmd *cobra.Command, args []string) error {
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
	results := as.Check().Fix()

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
