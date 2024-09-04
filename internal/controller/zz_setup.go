// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	environmentsecret "github.com/xunholy/provider-github/internal/controller/actions/environmentsecret"
	environmentvariable "github.com/xunholy/provider-github/internal/controller/actions/environmentvariable"
	organizationpermissions "github.com/xunholy/provider-github/internal/controller/actions/organizationpermissions"
	organizationvariable "github.com/xunholy/provider-github/internal/controller/actions/organizationvariable"
	repositoryaccesslevel "github.com/xunholy/provider-github/internal/controller/actions/repositoryaccesslevel"
	repositorypermissions "github.com/xunholy/provider-github/internal/controller/actions/repositorypermissions"
	secret "github.com/xunholy/provider-github/internal/controller/actions/secret"
	variable "github.com/xunholy/provider-github/internal/controller/actions/variable"
	organization "github.com/xunholy/provider-github/internal/controller/enterprise/organization"
	ruleset "github.com/xunholy/provider-github/internal/controller/enterprise/ruleset"
	membership "github.com/xunholy/provider-github/internal/controller/github/membership"
	providerconfig "github.com/xunholy/provider-github/internal/controller/providerconfig"
	branch "github.com/xunholy/provider-github/internal/controller/repo/branch"
	branchprotectionv3 "github.com/xunholy/provider-github/internal/controller/repo/branchprotectionv3"
	deploykey "github.com/xunholy/provider-github/internal/controller/repo/deploykey"
	file "github.com/xunholy/provider-github/internal/controller/repo/file"
	protection "github.com/xunholy/provider-github/internal/controller/repo/protection"
	pullrequest "github.com/xunholy/provider-github/internal/controller/repo/pullrequest"
	repository "github.com/xunholy/provider-github/internal/controller/repo/repository"
	repositoryautolinkreference "github.com/xunholy/provider-github/internal/controller/repo/repositoryautolinkreference"
	repositorywebhook "github.com/xunholy/provider-github/internal/controller/repo/repositorywebhook"
	rulesetrepo "github.com/xunholy/provider-github/internal/controller/repo/ruleset"
	emugroupmapping "github.com/xunholy/provider-github/internal/controller/team/emugroupmapping"
	members "github.com/xunholy/provider-github/internal/controller/team/members"
	team "github.com/xunholy/provider-github/internal/controller/team/team"
	teammembership "github.com/xunholy/provider-github/internal/controller/team/teammembership"
	teamrepository "github.com/xunholy/provider-github/internal/controller/team/teamrepository"
	teamsettings "github.com/xunholy/provider-github/internal/controller/team/teamsettings"
	teamsyncgroupmapping "github.com/xunholy/provider-github/internal/controller/team/teamsyncgroupmapping"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environmentsecret.Setup,
		environmentvariable.Setup,
		organizationpermissions.Setup,
		organizationvariable.Setup,
		repositoryaccesslevel.Setup,
		repositorypermissions.Setup,
		secret.Setup,
		variable.Setup,
		organization.Setup,
		ruleset.Setup,
		membership.Setup,
		providerconfig.Setup,
		branch.Setup,
		branchprotectionv3.Setup,
		deploykey.Setup,
		file.Setup,
		protection.Setup,
		pullrequest.Setup,
		repository.Setup,
		repositoryautolinkreference.Setup,
		repositorywebhook.Setup,
		repositorywebhook.Setup,
		rulesetrepo.Setup,
		emugroupmapping.Setup,
		members.Setup,
		team.Setup,
		teammembership.Setup,
		teamrepository.Setup,
		teamsettings.Setup,
		teamsyncgroupmapping.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
