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
	"get.porter.sh/porter/pkg/secrets"
	plugin "github.com/hashicorp/go-plugin"
)

// These are build-time values, set during an official release
var (
	Commit  string
	Version string
)

const VaultPluginInterface = secrets.PluginInterface + ".hashicorp.vault"

type PluginBox struct {
	*context.Context
	config.Config
}

func New() *PluginBox {
	return &PluginBox{
		Context: context.New(),
	}
}

func (p *PluginBox) Run(args []string) {
	if err := json.NewDecoder(p.In).Decode(&p.Config); err != nil {
		fmt.Fprint(p.Err, errors.Wrapf(err, "could not unmarshal config from input"))
	}

	key := args[0]
	parts := strings.Split(key, ".")
	// HANDLE index out of range
	selectedInterface := parts[0]
	var plugin plugin.Plugin
	switch key {
	case VaultPluginInterface:
		plugin = vault.NewPlugin(p.Config)
	}

	if plugin == nil {
		fmt.Fprintf(p.Err, "invalid plugin key specified: %q", key)
	}

	plugins.Serve(selectedInterface, plugin)

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
