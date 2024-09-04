package repositorytagprotection

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure github_repository_tag_protection resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_tag_protection", func(r *config.Resource) {
		r.Kind = "RepositoryWebhook"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
