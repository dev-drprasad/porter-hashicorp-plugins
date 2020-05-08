module github.com/dev-drprasad/porter-hashicorp-plugins

go 1.13

require (
	get.porter.sh/plugin/azure v0.2.0-beta.1
	get.porter.sh/porter v0.26.2-beta.1
	github.com/Azure/azure-sdk-for-go v19.1.1+incompatible
	github.com/cnabio/cnab-go v0.11.0-beta1
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/vault/api v1.0.4
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v1.0.0
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
