module github.com/dev-drprasad/porter-hashicorp-plugins

go 1.13

require (
	get.porter.sh/porter v0.26.2-beta.1
	github.com/spf13/cobra v1.0.0
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
