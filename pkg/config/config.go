package config

type Config struct {
	VaultAddr  string `json:"vault_addr"`
	VaultToken string `json:"vault_token"`
	PathPrefix string `json:"path_prefix"`
}
