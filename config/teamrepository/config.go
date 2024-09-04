package teamrepository

import "github.com/crossplane/upjet/pkg/config"

// Configure github_team_repository resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team_repository", func(r *config.Resource) {

		r.ShortGroup = "team"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
		r.References["team_id"] = config.Reference{
			TerraformName: "github_team",
		}
	})
}
