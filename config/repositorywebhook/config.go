package repositorywebhook

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure github_repository_webhook resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_webhook", func(r *config.Resource) {

		r.Kind = "RepositoryWebhook"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
	})
}
