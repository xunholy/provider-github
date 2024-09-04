package organizationruleset

import "github.com/crossplane/upjet/pkg/config"

// Configure github_organization_ruleset resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_organization_ruleset", func(r *config.Resource) {

		r.Kind = "OrganizationRuleset"
		r.ShortGroup = "enterprise"
	})
}
