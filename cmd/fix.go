package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/akerl/go-prospectus/checks"

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

	cs, err := checks.NewSet(params)
	if err != nil {
		return err
	}
	results := cs.Execute()

	fixResults := map[string]checks.ResultSet{
		"fixed":   {},
		"unfixed": {},
		"good":    {},
	}
	for _, item := range results {
		if item.Matches() {
			fixResults["good"] = append(fixResults["good"], item)
		} else {
			newResult := item.Fix()
			if newResult.Matches() {
				fixResults["fixed"] = append(fixResults["fixed"], newResult)
			} else {
				fixResults["unfixed"] = append(fixResults["unfixed"], newResult)
			}
		}
	}

	var output strings.Builder
	if flagJSON {
		outputBytes, err := json.MarshalIndent(fixResults, "", "  ")
		if err != nil {
			return err
		}
		output.Write(outputBytes)
	} else {
		for _, key := range []string{"good", "fixed", "unfixed"} {
			output.WriteString(fmt.Sprintf("%s:\n", key))
			for _, item := range fixResults[key] {
				output.WriteString(fmt.Sprintf("  %s\n", item))
			}
		}
	}
	fmt.Println(output.String())
	return nil
}
