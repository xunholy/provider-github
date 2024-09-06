package teamrepository

import "github.com/crossplane/upjet/pkg/config"

// Configure github_team_repository resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team_repository", func(r *config.Resource) {

		r.Kind = "TeamRepository"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
		r.References["team_id"] = config.Reference{
			TerraformName: "github_team",
		}
	})
}
