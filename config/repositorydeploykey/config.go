package repositorydeploykey

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure github_repository_deploy_key resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_deploy_key", func(r *config.Resource) {
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			Type: "github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository",
		}

		r.TerraformResource.Schema["key"].Required = true
		r.TerraformResource.Schema["read_only"].Required = true
		r.TerraformResource.Schema["title"].Required = true

		// Setting the field as sensitive to be able to pass the content from a k8s secret
		r.TerraformResource.Schema["key"].Sensitive = true
	})
}
