### Porter's HashiCorp Plugins

supports only `secrets` interface

Install:

```
porter plugin install hashicorp --feed-url https://github.com/dev-drprasad/porter-hashicorp-plugins/releases/download/feed/atom.xml
```

update `config.yml`:

```
default-secrets = "mysecrets"
[[secrets]]
name = "mysecrets"
plugin = "hashicorp.vault"

[secrets.config]
vault_addr = "http://vault.example.com:7500"
path_prefix = "kv"
vault_token = "token"
```
