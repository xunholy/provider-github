package enterpriseorganization

import "github.com/crossplane/upjet/pkg/config"

// Configure github_enterprise_organization resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_enterprise_organization", func(r *config.Resource) {

		r.ShortGroup = "enterprise"
	})
}
