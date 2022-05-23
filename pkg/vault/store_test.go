package vault

import (
	"context"
	"fmt"
	"testing"

	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ VaultClient = TestVaultClient{}

type TestVaultClient struct {
	Data map[string]map[string]interface{}
}

func NewTestVaultClient() TestVaultClient {
	return TestVaultClient{
		Data: map[string]map[string]interface{}{
			"data": map[string]interface{}{},
		},
	}
}

func (t TestVaultClient) Read(path string) (*vaultapi.Secret, error) {
	value, ok := t.Data[path]
	if !ok {
		return nil, fmt.Errorf("test secret %s was not mocked: %#v", path, t.Data)
	}

	return &vaultapi.Secret{Data: value}, nil
}

func (t TestVaultClient) Write(path string, value map[string]interface{}) (*vaultapi.Secret, error) {
	t.Data[path] = value
	return &vaultapi.Secret{Data: value}, nil
}

func TestStore_Resolve(t *testing.T) {
	cfg := config.Config{
		PathPrefix: "secret/myapp",
	}
	store := NewStore(cfg)
	testClient := NewTestVaultClient()
	store.client = testClient

	// Set up some test secret data
	testClient.Data["secret/data/myapp/v1"] = map[string]interface{}{
		"data": map[string]interface{}{
			"connstr": "myconnstr",
		},
	}

	got, err := store.Resolve(context.Background(), "secret", "v1/connstr")
	require.NoError(t, err, "Resolve failed")
	assert.Equal(t, "myconnstr", got, "incorrect secret resolved")
}

func TestStore_Create(t *testing.T) {
	cfg := config.Config{
		PathPrefix: "secret/myapp",
	}
	store := NewStore(cfg)
	testClient := NewTestVaultClient()
	store.client = testClient

	err := store.Create(context.Background(), "secret", "v1/connstr", "myconnstr")
	require.NoError(t, err, "Create failed")

	got, err := store.Resolve(context.Background(), "secret", "v1/connstr")
	require.NoError(t, err, "Resolve failed")
	assert.Equal(t, "myconnstr", got, "incorrect secret persisted")
}
