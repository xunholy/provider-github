// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	actionssecret "github.com/xunholy/provider-github/internal/controller/actions/actionssecret"
	actionsvariable "github.com/xunholy/provider-github/internal/controller/actions/actionsvariable"
	environmentsecret "github.com/xunholy/provider-github/internal/controller/actions/environmentsecret"
	environmentvariable "github.com/xunholy/provider-github/internal/controller/actions/environmentvariable"
	organizationpermissions "github.com/xunholy/provider-github/internal/controller/actions/organizationpermissions"
	organizationvariable "github.com/xunholy/provider-github/internal/controller/actions/organizationvariable"
	repositoryaccesslevel "github.com/xunholy/provider-github/internal/controller/actions/repositoryaccesslevel"
	repositorypermissions "github.com/xunholy/provider-github/internal/controller/actions/repositorypermissions"
	organization "github.com/xunholy/provider-github/internal/controller/enterprise/organization"
	organizationruleset "github.com/xunholy/provider-github/internal/controller/enterprise/organizationruleset"
	providerconfig "github.com/xunholy/provider-github/internal/controller/providerconfig"
	autolinkreference "github.com/xunholy/provider-github/internal/controller/repo/autolinkreference"
	branch "github.com/xunholy/provider-github/internal/controller/repo/branch"
	branchprotection "github.com/xunholy/provider-github/internal/controller/repo/branchprotection"
	branchprotectionv3 "github.com/xunholy/provider-github/internal/controller/repo/branchprotectionv3"
	defaultbranch "github.com/xunholy/provider-github/internal/controller/repo/defaultbranch"
	deploykey "github.com/xunholy/provider-github/internal/controller/repo/deploykey"
	environment "github.com/xunholy/provider-github/internal/controller/repo/environment"
	issuelabels "github.com/xunholy/provider-github/internal/controller/repo/issuelabels"
	pullrequest "github.com/xunholy/provider-github/internal/controller/repo/pullrequest"
	repository "github.com/xunholy/provider-github/internal/controller/repo/repository"
	repositoryfile "github.com/xunholy/provider-github/internal/controller/repo/repositoryfile"
	repositorywebhook "github.com/xunholy/provider-github/internal/controller/repo/repositorywebhook"
	ruleset "github.com/xunholy/provider-github/internal/controller/repo/ruleset"
	tagprotection "github.com/xunholy/provider-github/internal/controller/repo/tagprotection"
	groupmapping "github.com/xunholy/provider-github/internal/controller/team/groupmapping"
	members "github.com/xunholy/provider-github/internal/controller/team/members"
	membership "github.com/xunholy/provider-github/internal/controller/team/membership"
	settings "github.com/xunholy/provider-github/internal/controller/team/settings"
	syncgroupmapping "github.com/xunholy/provider-github/internal/controller/team/syncgroupmapping"
	team "github.com/xunholy/provider-github/internal/controller/team/team"
	teamrepository "github.com/xunholy/provider-github/internal/controller/team/teamrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		actionssecret.Setup,
		actionsvariable.Setup,
		environmentsecret.Setup,
		environmentvariable.Setup,
		organizationpermissions.Setup,
		organizationvariable.Setup,
		repositoryaccesslevel.Setup,
		repositorypermissions.Setup,
		organization.Setup,
		organizationruleset.Setup,
		providerconfig.Setup,
		autolinkreference.Setup,
		branch.Setup,
		branchprotection.Setup,
		branchprotectionv3.Setup,
		defaultbranch.Setup,
		deploykey.Setup,
		environment.Setup,
		issuelabels.Setup,
		pullrequest.Setup,
		repository.Setup,
		repositoryfile.Setup,
		repositorywebhook.Setup,
		ruleset.Setup,
		tagprotection.Setup,
		groupmapping.Setup,
		members.Setup,
		membership.Setup,
		membership.Setup,
		settings.Setup,
		syncgroupmapping.Setup,
		team.Setup,
		teamrepository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
