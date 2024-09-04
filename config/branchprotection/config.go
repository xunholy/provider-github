package branchprotection

import "github.com/crossplane/upjet/pkg/config"

// Configure github_branch_protection resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_branch_protection", func(r *config.Resource) {
		r.ShortGroup = "repo"

		r.References["repository_id"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
