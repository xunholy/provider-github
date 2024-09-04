package repositoryfile

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_file", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "repo"
		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}
		// r.References["branch"] = config.Reference{
		// 	Type: "github.com/xunholy/provider-github/apis/branch/v1alpha1.Branch",
		// }
	})
}
