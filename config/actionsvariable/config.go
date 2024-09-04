package actionsvariable

import "github.com/crossplane/upjet/pkg/config"

// Configure github_actions_variable resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_actions_variable", func(r *config.Resource) {

		r.Kind = "ActionsVariable"
		r.ShortGroup = "actions"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
