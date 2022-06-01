package pkg

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/vault"
	"github.com/pkg/errors"

	"get.porter.sh/porter/pkg/context"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/plugins"
	"get.porter.sh/porter/pkg/porter/version"
	plugin "github.com/hashicorp/go-plugin"
)

// These are build-time values, set during an official release
var (
	Commit  string
	Version string
)

const BuildPluginInterface = "build" + ".advanced"

type PluginBox struct {
	*context.Context
	config.Config
}

func New() *PluginBox {
	return &PluginBox{
		Context: context.New(),
	}
}

func (p *PluginBox) Run(args []string) error {
	if err := json.NewDecoder(p.In).Decode(&p.Config); err != nil {
		return errors.Wrapf(err, "could not unmarshal config from input")
	}

	var plugin plugin.Plugin
	key := args[0]
	switch key {
	case BuildPluginInterface:
		plugin = vault.NewPlugin(p.Config)
	}

	if plugin == nil {
		return errors.New(fmt.Sprintf("invalid plugin interface specified: %q", key))
	}

	parts := strings.Split(key, ".")
	selectedInterface := parts[0]
	plugins.Serve(selectedInterface, plugin)
	return nil
}

func (p *PluginBox) PrintVersion(opts version.Options) error {
	metadata := plugins.Metadata{
		Metadata: pkgmgmt.Metadata{
			Name: "hashicorp",
			VersionInfo: pkgmgmt.VersionInfo{
				Version: Version,
				Commit:  Commit,
				Author:  "REDDY PRASAD (@dev-drprasad)",
			},
		},
		Implementations: []plugins.Implementation{
			{Type: "secrets", Name: "vault"},
		},
	}

	return version.PrintVersion(p.Context, opts, metadata)
}
