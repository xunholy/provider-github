package repositoryruleset

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_ruleset resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_ruleset", func(r *config.Resource) {

		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
