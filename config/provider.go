/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/xunholy/provider-github/config/actionsenvironmentsecret"
	"github.com/xunholy/provider-github/config/actionsenvironmentvariable"
	"github.com/xunholy/provider-github/config/actionsorganizationpermissions"
	"github.com/xunholy/provider-github/config/actionsorganizationvariable"
	"github.com/xunholy/provider-github/config/actionsrepositoryaccesslevel"
	"github.com/xunholy/provider-github/config/actionsrepositorypermissions"
	"github.com/xunholy/provider-github/config/actionssecret"
	"github.com/xunholy/provider-github/config/actionsvariable"
	"github.com/xunholy/provider-github/config/branch"
	"github.com/xunholy/provider-github/config/branchdefault"
	"github.com/xunholy/provider-github/config/branchprotection"
	"github.com/xunholy/provider-github/config/branchprotectionv3"
	"github.com/xunholy/provider-github/config/emugroupmapping"
	"github.com/xunholy/provider-github/config/enterpriseorganization"
	"github.com/xunholy/provider-github/config/issuelabels"
	"github.com/xunholy/provider-github/config/membership"
	"github.com/xunholy/provider-github/config/organizationruleset"
	"github.com/xunholy/provider-github/config/repository"
	"github.com/xunholy/provider-github/config/repositoryautolinkreference"
	"github.com/xunholy/provider-github/config/repositorydeploykey"
	"github.com/xunholy/provider-github/config/repositoryenvironment"
	"github.com/xunholy/provider-github/config/repositoryfile"
	"github.com/xunholy/provider-github/config/repositorypullrequest"
	"github.com/xunholy/provider-github/config/repositoryruleset"
	"github.com/xunholy/provider-github/config/repositorytagprotection"
	"github.com/xunholy/provider-github/config/repositorywebhook"
	"github.com/xunholy/provider-github/config/team"
	"github.com/xunholy/provider-github/config/teammembers"
	"github.com/xunholy/provider-github/config/teammembership"
	"github.com/xunholy/provider-github/config/teamrepository"
	"github.com/xunholy/provider-github/config/teamsettings"
	"github.com/xunholy/provider-github/config/teamsyncgroupmapping"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/xunholy/provider-github"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		actionsenvironmentsecret.Configure,
		actionsenvironmentvariable.Configure,
		actionsorganizationpermissions.Configure,
		actionsorganizationvariable.Configure,
		actionsrepositoryaccesslevel.Configure,
		actionsrepositorypermissions.Configure,
		actionssecret.Configure,
		actionsvariable.Configure,
		branch.Configure,
		branchdefault.Configure,
		branchprotection.Configure,
		branchprotectionv3.Configure,
		emugroupmapping.Configure,
		enterpriseorganization.Configure,
		issuelabels.Configure,
		membership.Configure,
		organizationruleset.Configure,
		repository.Configure,
		repositoryautolinkreference.Configure,
		repositorydeploykey.Configure,
		repositoryenvironment.Configure,
		repositoryfile.Configure,
		repositorypullrequest.Configure,
		repositoryruleset.Configure,
		repositorytagprotection.Configure,
		repositorywebhook.Configure,
		team.Configure,
		teammembers.Configure,
		teammembership.Configure,
		teamrepository.Configure,
		teamsettings.Configure,
		teamsyncgroupmapping.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
