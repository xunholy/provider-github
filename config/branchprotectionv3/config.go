package branchprotectionv3

import "github.com/crossplane/upjet/pkg/config"

// Configure github_branch_protection_v3 resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_branch_protection_v3", func(r *config.Resource) {
		r.Kind = "BranchProtectionV3"
		r.ShortGroup = "repo"

		r.References["repository_id"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
