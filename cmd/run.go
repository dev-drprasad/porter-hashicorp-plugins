package cmd

import (
	"fmt"
	"os"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:       "run [implementation]",
	Short:     "Run the plugin and listen for client connections.",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{pkg.VaultPluginInterface},
	Run: func(cmd *cobra.Command, args []string) {
		p.Run(args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		si, err := os.Stdin.Stat()
		if err != nil {
			fmt.Fprint(p.Err, "could not get stdin info")
			os.Exit(1)
		}
		if (si.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			fmt.Fprintf(p.Err, "This binary is a Porter plugin. It is not meant to be executed directly.")
			os.Exit(126)
		}
	},
}
