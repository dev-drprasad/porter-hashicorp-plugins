package vault

import (
	"get.porter.sh/porter/pkg/secrets"
	cnabsecrets "github.com/cnabio/cnab-go/secrets"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	plugin "github.com/hashicorp/go-plugin"
)

var _ cnabsecrets.Store = &Plugin{}

// Plugin is the plugin wrapper for accessing secrets from Vault.
type Plugin struct {
	cnabsecrets.Store
}

func NewPlugin(cfg config.Config) plugin.Plugin {
	return &secrets.Plugin{
		Impl: &Plugin{
			Store: NewStore(cfg),
		},
	}
}
