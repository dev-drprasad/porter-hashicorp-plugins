package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/plugins"
	"get.porter.sh/porter/pkg/porter/version"
	"get.porter.sh/porter/pkg/portercontext"
	secretplugins "get.porter.sh/porter/pkg/secrets/plugins"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/vault"
	"github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
)

// These are build-time values, set during an official release
var (
	Commit  string
	Version string
)

const VaultPluginInterface = secretplugins.PluginInterface + ".hashicorp.vault"

type PluginBox struct {
	*portercontext.Context
	config.Config
}

func New() *PluginBox {
	return &PluginBox{
		Context: portercontext.New(),
	}
}

func (p *PluginBox) Run(ctx context.Context, args []string) error {
	if err := json.NewDecoder(p.In).Decode(&p.Config); err != nil {
		return errors.Wrapf(err, "could not unmarshal config from input")
	}

	var plugin plugin.Plugin
	key := args[0]
	switch key {
	case VaultPluginInterface:
		plugin = vault.NewPlugin(p.Context, p.Config)
	}

	if plugin == nil {
		return errors.New(fmt.Sprintf("invalid plugin interface specified: %q", key))
	}

	parts := strings.Split(key, ".")
	selectedInterface := parts[0]
	plugins.Serve(p.Context, selectedInterface, plugin, secretplugins.PluginProtocolVersion)
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
