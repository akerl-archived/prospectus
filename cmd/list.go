package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/akerl/prospectus/checks"

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

	cs, err := checks.NewSet(params)
	if err != nil {
		return err
	}
	if cs == nil {
		cs = checks.CheckSet{}
	}

	var output strings.Builder
	if flagJSON {
		outputBytes, err := json.MarshalIndent(cs, "", "  ")
		if err != nil {
			return err
		}
		output.Write(outputBytes)
	} else {
		for _, item := range cs {
			output.WriteString(item.String())
			output.WriteString("\n")
		}
	}
	fmt.Println(output.String())
	return nil
}
