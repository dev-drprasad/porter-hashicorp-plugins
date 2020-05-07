package cmd

import (
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg"
	"github.com/spf13/cobra"
)

var p *pkg.PluginBox

var rootCmd = &cobra.Command{
	Use:   "hashicorp",
	Short: "HashiCorp plugins for Porter",
}

func init() {
	p = pkg.New()
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
