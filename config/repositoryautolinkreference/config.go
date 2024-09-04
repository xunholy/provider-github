package repositoryautolinkreference

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_autolink_reference resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_autolink_reference", func(r *config.Resource) {

		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
	})
}
