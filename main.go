package main

import (
	"fmt"
	"os"

	"github.com/akerl/prospectus-ng/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
