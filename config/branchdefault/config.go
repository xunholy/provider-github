package branchdefault

import "github.com/crossplane/upjet/pkg/config"

// Configure github_branch_default resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_branch_default", func(r *config.Resource) {

		r.Kind = "DefaultBranch"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
		r.References["branch"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Branch",
		}
	})
}
