package repositoryruleset

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_ruleset resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_ruleset", func(r *config.Resource) {

		r.Kind = "RepositoryRuleset"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
	})
}
