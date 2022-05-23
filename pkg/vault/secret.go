package vault

import (
	"fmt"
	"path"
	"strings"
)

// Secret represents the path to a secret in vault.
type Secret struct {
	// Path for retrieving the secret's data.
	readPath string

	// Path for saving the secret's data.
	// Data is always written to the Config.PorterSecret sub-path.
	writePath string

	// Path displayed to the user.
	userPath string

	// Key for retrieving the secret value from the Vault secret.
	key string
}

func (p Secret) ReadPath() string {
	return p.readPath
}

func (p Secret) WritePath() string {
	return p.writePath
}

func (p Secret) SecretPath() string {
	return p.userPath
}

func (p Secret) String() string {
	return fmt.Sprintf("%s[%s]", p.userPath, p.key)
}

// NewSecret constructs the path to the specified secret using
// the plugin configuration.
func (s *Store) NewSecret(keyValue string) Secret {
	secret := Secret{}

	// Secrets are all resolved from the configured path prefix, e.g. mysecrets,
	// then the secret defined in the ParameterSet or CredentialSet may include additional subpaths, e.g. myapp/v1/connstr
	// the secret is defined at mysecrets/data/myapp/v1, and the final part of the path is used as the key for the secret data (i.e. the key will be connstr)
	keyPrefix := path.Dir(keyValue)
	if keyPrefix == keyValue {
		keyPrefix = ""
	}
	secret.key = path.Base(keyValue)
	engine, readPathSuffix, _ := strings.Cut(s.config.PathPrefix, "/")

	// path for retrieving data from the secret, has a hard coded /data/ segment that users don't see what using vault directly
	secret.readPath = path.Join(engine, "data", readPathSuffix, keyPrefix)

	// path for saving data to the secret, stores it as a key under the porter secret
	secret.writePath = path.Join(engine, "data", readPathSuffix, s.config.GetPorterSecret(), keyPrefix)

	// Secret is the path that we should show to the user when there's a problem
	secret.userPath = path.Join(engine, readPathSuffix, keyPrefix)

	return secret
}
