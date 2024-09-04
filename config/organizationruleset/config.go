package organizationruleset

import "github.com/crossplane/upjet/pkg/config"

// Configure github_organization_ruleset resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_organization_ruleset", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.Kind = "OrganizationRuleset"
		r.ShortGroup = "enterprise"
	})
}
