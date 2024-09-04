package repositoryfile

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_file resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_file", func(r *config.Resource) {

		r.Kind = "RepositoryFile"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
		r.References["branch"] = config.Reference{
			TerraformName: "github_branch",
		}
	})
}
