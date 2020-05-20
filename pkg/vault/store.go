package vault

import (
	"encoding/json"
	"fmt"
	"strings"

	"get.porter.sh/porter/pkg/secrets"
	cnabsecrets "github.com/cnabio/cnab-go/secrets"
	"github.com/dev-drprasad/porter-hashicorp-plugins/pkg/config"
	"github.com/hashicorp/vault/api"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

var _ cnabsecrets.Store = &Store{}

const (
	SecretKeyName = "secret"
)

// Store implements the backing store for secrets in azure key vault.
type Store struct {
	config config.Config
	client *vaultapi.Client
}

func NewStore(cfg config.Config) cnabsecrets.Store {
	s := &Store{
		config: cfg,
	}

	return secrets.NewSecretStore(s)
}

func (s *Store) Connect() error {
	if s.client != nil {
		return nil
	}

	config := &vaultapi.Config{
		Address: s.config.VaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		return errors.Wrapf(err, "could not connect to vault server with address %s", s.config.VaultAddr)
	}
	s.client = client
	s.client.SetToken(s.config.VaultToken)

	return nil
}

func (s *Store) Resolve(keyName string, keyValue string) (string, error) {
	if strings.ToLower(keyName) != SecretKeyName {
		return "", errors.Errorf("could not resolve unsupported keyName '%s'. Vault plugin only supports '%s' right now", keyName, SecretKeyName)
	}

	vaultSecret, err := s.client.Logical().Read(s.config.PathPrefix + "/data/" + keyValue)
	if err != nil {
		return "", errors.Wrapf(err, "error while reading \"%s\" from vault", keyValue)
	}
	if vaultSecret == nil {
		return "", errors.New(fmt.Sprintf("no secret value found at \"%s\"", keyValue))
	}

	data, ok := vaultSecret.Data["data"]
	if !ok {
		return "", errors.New("property 'data' does not exist in secret")
	}

	secretB, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "could not marshal secret data")
	}

	return string(secretB), nil
}
