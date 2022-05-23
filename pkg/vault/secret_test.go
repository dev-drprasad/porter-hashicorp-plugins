package vault

import (
	"testing"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSecret(t *testing.T) {
	cfg := config.Config{
		PathPrefix:   "secret/myapp",
		PorterSecret: "myporter",
	}
	store := NewStore(cfg)

	secret := store.NewSecret("v1/connstr")

	assert.Equal(t, "secret/myapp/v1", secret.SecretPath(), "incorrect display path")
	assert.Equal(t, "secret/data/myapp/v1", secret.ReadPath(), "incorrect read path")
	assert.Equal(t, "secret/data/myapp/myporter/v1", secret.WritePath(), "incorrect write path")
	assert.Equal(t, "secret/myapp/v1[connstr]", secret.String(), "incorrect string representation")
}
