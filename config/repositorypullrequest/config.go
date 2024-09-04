package repositorypullrequest

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_pull_request resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_pull_request", func(r *config.Resource) {

		r.ShortGroup = "repo"

		r.References["base_repository"] = config.Reference{
			TerraformName: "github_repository",
		}
		//    r.References["base_ref"] = config.Reference{
		//			Type: "github.com/coopnorge/provider-github/apis/repo/v1alpha1.Branch",
		//		}
		r.References["head_ref"] = config.Reference{
			TerraformName: "github_branch",
		}

	})
}
