// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AutolinkReference.
func (mg *AutolinkReference) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Branch.
func (mg *Branch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BranchProtection.
func (mg *BranchProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryIDRef,
		Selector:     mg.Spec.ForProvider.RepositoryIDSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryID")
	}
	mg.Spec.ForProvider.RepositoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RepositoryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryIDRef,
		Selector:     mg.Spec.InitProvider.RepositoryIDSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RepositoryID")
	}
	mg.Spec.InitProvider.RepositoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BranchProtectionv3.
func (mg *BranchProtectionv3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultBranch.
func (mg *DefaultBranch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BranchRef,
		Selector:     mg.Spec.ForProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Branch")
	}
	mg.Spec.ForProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BranchRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BranchRef,
		Selector:     mg.Spec.InitProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Branch")
	}
	mg.Spec.InitProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BranchRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeployKey.
func (mg *DeployKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Reviewers); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Reviewers[i3].Teams),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Reviewers[i3].TeamsRefs,
			Selector:      mg.Spec.ForProvider.Reviewers[i3].TeamsSelector,
			To: reference.To{
				List:    &github_teamList{},
				Managed: &github_team{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Reviewers[i3].Teams")
		}
		mg.Spec.ForProvider.Reviewers[i3].Teams = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Reviewers[i3].TeamsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Reviewers); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Reviewers[i3].Users),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Reviewers[i3].UsersRefs,
			Selector:      mg.Spec.ForProvider.Reviewers[i3].UsersSelector,
			To: reference.To{
				List:    &github_userList{},
				Managed: &github_user{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Reviewers[i3].Users")
		}
		mg.Spec.ForProvider.Reviewers[i3].Users = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Reviewers[i3].UsersRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Reviewers); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Reviewers[i3].Teams),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.Reviewers[i3].TeamsRefs,
			Selector:      mg.Spec.InitProvider.Reviewers[i3].TeamsSelector,
			To: reference.To{
				List:    &github_teamList{},
				Managed: &github_team{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Reviewers[i3].Teams")
		}
		mg.Spec.InitProvider.Reviewers[i3].Teams = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Reviewers[i3].TeamsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Reviewers); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Reviewers[i3].Users),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.Reviewers[i3].UsersRefs,
			Selector:      mg.Spec.InitProvider.Reviewers[i3].UsersSelector,
			To: reference.To{
				List:    &github_userList{},
				Managed: &github_user{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Reviewers[i3].Users")
		}
		mg.Spec.InitProvider.Reviewers[i3].Users = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Reviewers[i3].UsersRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this EnvironmentDeploymentPolicy.
func (mg *EnvironmentDeploymentPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Environment),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvironmentRef,
		Selector:     mg.Spec.ForProvider.EnvironmentSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Environment")
	}
	mg.Spec.ForProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Environment),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.EnvironmentRef,
		Selector:     mg.Spec.InitProvider.EnvironmentSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Environment")
	}
	mg.Spec.InitProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvironmentRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PullRequest.
func (mg *PullRequest) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaseRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BaseRefRef,
		Selector:     mg.Spec.ForProvider.BaseRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaseRef")
	}
	mg.Spec.ForProvider.BaseRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaseRefRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaseRepository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BaseRepositoryRef,
		Selector:     mg.Spec.ForProvider.BaseRepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaseRepository")
	}
	mg.Spec.ForProvider.BaseRepository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaseRepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HeadRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HeadRefRef,
		Selector:     mg.Spec.ForProvider.HeadRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HeadRef")
	}
	mg.Spec.ForProvider.HeadRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HeadRefRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BaseRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BaseRefRef,
		Selector:     mg.Spec.InitProvider.BaseRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BaseRef")
	}
	mg.Spec.InitProvider.BaseRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BaseRefRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BaseRepository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BaseRepositoryRef,
		Selector:     mg.Spec.InitProvider.BaseRepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BaseRepository")
	}
	mg.Spec.InitProvider.BaseRepository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BaseRepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HeadRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HeadRefRef,
		Selector:     mg.Spec.InitProvider.HeadRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HeadRef")
	}
	mg.Spec.InitProvider.HeadRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HeadRefRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryFile.
func (mg *RepositoryFile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BranchRef,
		Selector:     mg.Spec.ForProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Branch")
	}
	mg.Spec.ForProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BranchRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Branch),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BranchRef,
		Selector:     mg.Spec.InitProvider.BranchSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Branch")
	}
	mg.Spec.InitProvider.Branch = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BranchRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryRuleset.
func (mg *RepositoryRuleset) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryWebhook.
func (mg *RepositoryWebhook) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagProtection.
func (mg *TagProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
