/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"github_actions_environment_secret":               config.IdentifierFromProvider,
	"github_actions_environment_variable":             config.IdentifierFromProvider,
	"github_actions_organization_permissions":         config.IdentifierFromProvider,
	"github_actions_organization_variable":            config.IdentifierFromProvider,
	"github_actions_repository_access_level":          config.IdentifierFromProvider,
	"github_actions_repository_permissions":           config.IdentifierFromProvider,
	"github_actions_secret":                           config.IdentifierFromProvider,
	"github_actions_variable":                         config.IdentifierFromProvider,
	"github_branch":                                   config.IdentifierFromProvider,
	"github_branch_default":                           config.TemplatedStringAsIdentifier("repository", "{{ .external_name }}"),
	"github_branch_protection":                        config.IdentifierFromProvider,
	"github_branch_protection_v3":                     config.IdentifierFromProvider,
	"github_emu_group_mapping":                        config.IdentifierFromProvider,
	"github_enterprise_organization":                  config.IdentifierFromProvider,
	"github_issue_labels":                             config.IdentifierFromProvider,
	"github_membership":                               config.IdentifierFromProvider,
	"github_organization_ruleset":                     config.IdentifierFromProvider,
	"github_repository":                               config.IdentifierFromProvider,
	"github_repository_autolink_reference":            config.IdentifierFromProvider,
	"github_repository_deploy_key":                    config.IdentifierFromProvider,
	"github_repository_environment":                   config.IdentifierFromProvider,
	"github_repository_environment_deployment_policy": config.IdentifierFromProvider,
	"github_repository_file":                          config.IdentifierFromProvider,
	"github_repository_pull_request":                  config.IdentifierFromProvider,
	"github_repository_ruleset":                       config.IdentifierFromProvider,
	"github_repository_tag_protection":                config.IdentifierFromProvider,
	"github_repository_webhook":                       config.IdentifierFromProvider,
	"github_team":                                     config.IdentifierFromProvider,
	"github_team_members":                             config.IdentifierFromProvider,
	"github_team_membership":                          config.IdentifierFromProvider,
	"github_team_repository":                          config.IdentifierFromProvider,
	"github_team_settings":                            config.IdentifierFromProvider,
	"github_team_sync_group_mapping":                  config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
