package main

import (
	"os"

	"github.com/richbai90/porter-advanced-build-plugins/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
