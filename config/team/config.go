package team

import "github.com/crossplane/upjet/pkg/config"

// Configure github_team resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team", func(r *config.Resource) {
		r.ShortGroup = "team"
	})
}
