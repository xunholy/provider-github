package repositoryenvironment

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure github_repository_environment resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_environment", func(r *config.Resource) {

		r.ShortGroup = "repo"

		// Reference for the repository
		r.References["repository"] = config.Reference{
			Type: "github_repository",
		}

		// Reference for teams in the reviewers block
		r.References["reviewers.teams"] = config.Reference{
			Type: "github_team",
		}

		// Reference for users in the reviewers block
		r.References["reviewers.users"] = config.Reference{
			Type: "github_user",
		}

		// Configure the reviewers block
		r.TerraformResource.Schema["reviewers"].Elem.(*schema.Resource).Schema["teams"].Elem = &schema.Schema{
			Type: schema.TypeString,
		}
		r.TerraformResource.Schema["reviewers"].Elem.(*schema.Resource).Schema["users"].Elem = &schema.Schema{
			Type: schema.TypeString,
		}

		// Make sure the reviewers block is correctly handled
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"reviewers"},
		}
	})
}
