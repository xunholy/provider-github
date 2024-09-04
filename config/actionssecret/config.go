package actionssecret

import "github.com/crossplane/upjet/pkg/config"

// Configure github_actions_secret resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_actions_secret", func(r *config.Resource) {

		r.Kind = "ActionsSecret"
		r.ShortGroup = "actions"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
