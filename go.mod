module github.com/richbai90/porter-advanced-build-plugins

go 1.13

require (
	get.porter.sh/porter v0.26.2-beta.1
	github.com/cnabio/cnab-go v0.11.0-beta1
	github.com/docker/cli v0.0.0-20191017083524-a8ff7f821017
	github.com/docker/docker v1.4.2-0.20181229214054-f76d6a078d88
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/vault/api v1.0.4
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v1.0.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/genproto v0.0.0-20191007204434-a023cd5227bd // indirect
	google.golang.org/grpc v1.24.0 // indirect
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
