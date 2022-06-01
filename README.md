### Porter's HashiCorp Plugins

This is a [Porter][porter] plugin to inject credentials to Porter bundle from hashicorp's [vault][vault].

Supports porter version greater or equal to **v0.23.0-beta.1** and supports only [KV Version 2][kv-v2] secret engine. Please raise an issue if you're looking for support for other secret engines

#### Install:

```
porter plugin install hashicorp --feed-url https://github.com/richbai90/porter-advanced-build-plugins/releases/download/feed/atom.xml
```

#### Configuration:

To use vault plugin, add the following config to porter's config file (default location: `~/.porter/config.toml`).
Replace `vault_addr`, `vault_token` and `path_prefix` with proper values.

```
default-secrets = "porter-secrets"
[[secrets]]
name = "porter-secrets"
plugin = "hashicorp.vault"

[secrets.config]
vault_addr = "http://vault.example.com:7500"
path_prefix = "organization/team/project"
vault_token = "token"
```

Your secret will be injected as json into porter manifest. Currently there is no support for accessing specific key from a secret. If you're looking for that feature, do raise an issue.

#### Config Parameters

##### `path_prefix`

`path_prefix` lets allow you to specify prefix for your secret path. Let' say you have a secret (`myawesomeproject`) with path `organization/team/project/myawesomeproject`, then you can configure `path_prefix` as `organization/team/project`.

[porter]: https://porter.sh/
[vault]: https://www.vaultproject.io/
[kv-v2]: https://www.vaultproject.io/api-docs/secret/kv/kv-v2
