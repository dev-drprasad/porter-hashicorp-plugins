package config

const (
	// DefaultPorterSecret is the default location where porter's data is stored.
	DefaultPorterSecret = "porter"
)

type Config struct {
	VaultAddr  string `json:"vault_addr"`
	VaultToken string `json:"vault_token"`
	PathPrefix string `json:"path_prefix"`

	// PorterSecret is the path to the secret where Porter should store any sensitive
	// data that it generates, for example a sensitive bundle output.
	// The full path of the secret is PathPrefix + PorterSecret.
	PorterSecret string `json:"porter_secret"`
}

// GetPorterSecret returns the path of the secret where Porter can store secrets
// that it generates.
func (c Config) GetPorterSecret() string {
	if c.PorterSecret != "" {
		return c.PorterSecret
	}

	return DefaultPorterSecret
}
