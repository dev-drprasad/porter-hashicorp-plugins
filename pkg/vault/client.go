package vault

import (
	vaultapi "github.com/hashicorp/vault/api"
)

// VaultClient is a wrapper around the Vault API Client
// so that we can write unit tests
type VaultClient interface {
	Read(path string) (*vaultapi.Secret, error)
	Write(path string, data map[string]interface{}) (*vaultapi.Secret, error)
}
