package branch

import "github.com/crossplane/upjet/pkg/config"

// Configure github_branch resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_branch", func(r *config.Resource) {

		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
	})
}
