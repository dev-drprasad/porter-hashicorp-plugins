package vault

import (
	"context"
	"fmt"
	"strings"

	secretplugins "get.porter.sh/porter/pkg/secrets/plugins"
	"get.porter.sh/porter/pkg/secrets/plugins/host"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

var _ secretplugins.SecretsProtocol = &Store{}

const (
	SecretKeyName = "secret"
)

// Store implements the backing store for secrets in azure key vault.
type Store struct {
	config    config.Config
	client    VaultClient
	hostStore host.Store
}

func NewStore(cfg config.Config) *Store {
	s := &Store{
		config:    cfg,
		hostStore: host.NewStore(),
	}

	return s
}

func (s *Store) Connect() error {
	if s.client != nil {
		return nil
	}

	config := &vaultapi.Config{
		Address: s.config.VaultAddr,
	}
	client, err := vaultapi.NewClient(config)
	if err != nil {
		return errors.Wrapf(err, "could not connect to vault server with address %s", s.config.VaultAddr)
	}
	client.SetToken(s.config.VaultToken)
	s.client = client.Logical()

	return nil
}

func (s *Store) Resolve(ctx context.Context, keyName string, keyValue string) (string, error) {
	if err := s.Connect(); err != nil {
		return "", err
	}

	// If it isn't a secret, have the host plugin resolve it (which can handle env, command, value, etc)
	if strings.ToLower(keyName) != SecretKeyName {
		return s.hostStore.Resolve(ctx, keyName, keyValue)
	}

	secret := s.NewSecret(keyValue)

	vaultSecret, err := s.client.Read(secret.ReadPath())
	if err != nil {
		return "", errors.Wrapf(err, "error while reading %s from vault", secret.SecretPath())
	}
	if vaultSecret == nil {
		return "", errors.New(fmt.Sprintf("no secret found at %s in vault", secret.SecretPath()))
	}

	secretData, ok := vaultSecret.Data["data"].(map[string]interface{})
	if !ok {
		return "", errors.Errorf("property 'data' does not exist in secret %s", secret.SecretPath())
	}

	secretValue, ok := secretData[secret.key]
	if !ok {
		return "", errors.Errorf("key '%s' does not existing in secret '%s'", secret.key, secret.SecretPath())
	}
	return fmt.Sprintf("%v", secretValue), nil
}

func (s *Store) Create(ctx context.Context, keyName string, keyValue string, value string) error {
	if err := s.Connect(); err != nil {
		return err
	}

	secret := s.NewSecret(keyValue)
	_, err := s.client.Write(secret.ReadPath(), map[string]interface{}{
		"data": map[string]interface{}{
			secret.key: value,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating secret %s: %w", secret, err)
	}

	return nil
}
