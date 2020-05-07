package pkg

import (
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"

	"get.porter.sh/porter/pkg/context"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
)

// These are build-time values, set during an official release
var (
	Commit  string
	Version string
)

type PluginBox struct {
	*context.Context
	config.Config
}

func New() *PluginBox {
	return &PluginBox{
		Context: context.New(),
	}
}

func (p *PluginBox) PrintVersion(opts version.Options) error {
	metadata := pkgmgmt.Metadata{
		Name: "hashicorp",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: Version,
			Commit:  Commit,
			Author:  "REDDY PRASAD (@dev-drprasad)",
		},
	}

	return version.PrintVersion(p.Context, opts, metadata)
}
