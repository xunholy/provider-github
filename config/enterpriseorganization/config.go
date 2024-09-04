package enterpriseorganization

import "github.com/crossplane/upjet/pkg/config"

// Configure github_enterprise_organization resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_enterprise_organization", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "enterprise"
	})
}
