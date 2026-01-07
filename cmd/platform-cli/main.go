package main

import (
	"fmt"
	"os"

	"github.com/abevz/platform-iac-cli/internal/cli"
)

func main() {
	// Execute the root command from internal/cli package
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
