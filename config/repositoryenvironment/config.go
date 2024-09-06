package repositoryenvironment

import (
	"github.com/crossplane/upjet/pkg/config"
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

		// Make sure the reviewers block is correctly handled
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"reviewers"},
		}
	})
}
