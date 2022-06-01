package cmd

import (
	"errors"
	"os"

	"github.com/richbai90/porter-advanced-build-plugins/pkg"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:       "run [implementation]",
	Short:     "Run the plugin and listen for client connections.",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{pkg.BuildPluginInterface},
	RunE: func(cmd *cobra.Command, args []string) error {
		return p.Run(args)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		si, err := os.Stdin.Stat()
		if err != nil {
			return errors.New("could not get stdin info")
		}
		if (si.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			return errors.New("this binary is a Porter plugin. It is not meant to be executed directly")
		}
		return nil
	},
}
