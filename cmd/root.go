package cmd

import (
	"fmt"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg"
	"github.com/spf13/cobra"
)

var p *pkg.PluginBox

var rootCmd = &cobra.Command{
	Use:   "hashicorp",
	Short: "HashiCorp plugins for Porter",

	// Lets not send usage text and "Run ... for help" to caller when error happens
	// https://github.com/spf13/cobra/issues/1111
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	p = pkg.New()
	rootCmd.SetErr(p.Err)
	rootCmd.AddCommand(versionCmd, runCmd)
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(p.Err, err)
	}
	return err
}
