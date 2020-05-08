package cmd

import (
	"get.porter.sh/porter/pkg/porter/version"
	"github.com/spf13/cobra"
)

var versionFormat string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the plugin version",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := version.Options{}
		opts.RawFormat = versionFormat
		if err := opts.Validate(); err != nil {
			return err
		}
		return p.PrintVersion(opts)
	},
}

func init() {
	versionCmd.Flags().StringVarP(&versionFormat, "output", "o", string(version.DefaultVersionFormat), "Specify an output format.  Allowed values: json, plaintext")
}
