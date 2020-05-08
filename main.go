package main

import (
	"os"

	"github.com/dev-drprasad/porter-hashicorp-plugins/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
