package membership

import "github.com/crossplane/upjet/pkg/config"

// Configure github_membership resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_membership", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		// r.ShortGroup = "enterprise"
	})
}
