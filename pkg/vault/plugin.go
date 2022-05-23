package vault

import (
	"get.porter.sh/porter/pkg/portercontext"
	"get.porter.sh/porter/pkg/secrets/pluginstore"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	"github.com/hashicorp/go-plugin"
)

func NewPlugin(c *portercontext.Context, pluginCfg config.Config) plugin.Plugin {
	return pluginstore.NewPlugin(c, NewStore(pluginCfg))
}
