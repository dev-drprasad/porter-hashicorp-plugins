package cmd

import (
	"fmt"
	"os"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg"
	"github.com/spf13/cobra"
)

var p *pkg.PluginBox

var rootCmd = &cobra.Command{
	Use:   "hashicorp",
	Short: "HashiCorp plugins for Porter",
}

func init() {
	si, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("could not get stdin info")
		os.Exit(1)
	}
	if si.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("This binary is a Porter plugin. It is not meant to be executed directly.")
		os.Exit(126)
	}
	p = pkg.New()
	rootCmd.SetErr(p.Err)
	rootCmd.AddCommand(versionCmd, runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
