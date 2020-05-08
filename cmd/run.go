package cmd

import (
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:          "run [implementation]",
	Short:        "Run the plugin and listen for client connections.",
	Args:         cobra.ExactValidArgs(1),
	ValidArgs:    []string{pkg.VaultPluginInterface},
	SilenceUsage: true, // Lets not send usage text to caller when error happens
	Run: func(cmd *cobra.Command, args []string) {
		p.Run(args)
	},
}
