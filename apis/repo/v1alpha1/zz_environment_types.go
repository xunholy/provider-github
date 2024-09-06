// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeploymentBranchPolicyInitParameters struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	CustomBranchPolicies *bool `json:"customBranchPolicies,omitempty" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	ProtectedBranches *bool `json:"protectedBranches,omitempty" tf:"protected_branches,omitempty"`
}

type DeploymentBranchPolicyObservation struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	CustomBranchPolicies *bool `json:"customBranchPolicies,omitempty" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	ProtectedBranches *bool `json:"protectedBranches,omitempty" tf:"protected_branches,omitempty"`
}

type DeploymentBranchPolicyParameters struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	// +kubebuilder:validation:Optional
	CustomBranchPolicies *bool `json:"customBranchPolicies" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	// +kubebuilder:validation:Optional
	ProtectedBranches *bool `json:"protectedBranches" tf:"protected_branches,omitempty"`
}

type EnvironmentInitParameters struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	DeploymentBranchPolicy []DeploymentBranchPolicyInitParameters `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The environment reviewers configuration.
	Reviewers []ReviewersInitParameters `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *float64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type EnvironmentObservation struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	DeploymentBranchPolicy []DeploymentBranchPolicyObservation `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The environment reviewers configuration.
	Reviewers []ReviewersObservation `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *float64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type EnvironmentParameters struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	// +kubebuilder:validation:Optional
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	// +kubebuilder:validation:Optional
	DeploymentBranchPolicy []DeploymentBranchPolicyParameters `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	// +kubebuilder:validation:Optional
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The environment reviewers configuration.
	// +kubebuilder:validation:Optional
	Reviewers []ReviewersParameters `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	// +kubebuilder:validation:Optional
	WaitTimer *float64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type ReviewersInitParameters struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github_team
	// +listType=set
	Teams []*float64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// References to github_team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsRefs []v1.Reference `json:"teamsRefs,omitempty" tf:"-"`

	// Selector for a list of github_team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsSelector *v1.Selector `json:"teamsSelector,omitempty" tf:"-"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github_user
	// +listType=set
	Users []*float64 `json:"users,omitempty" tf:"users,omitempty"`

	// References to github_user to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of github_user to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

type ReviewersObservation struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +listType=set
	Teams []*float64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +listType=set
	Users []*float64 `json:"users,omitempty" tf:"users,omitempty"`
}

type ReviewersParameters struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github_team
	// +kubebuilder:validation:Optional
	// +listType=set
	Teams []*float64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// References to github_team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsRefs []v1.Reference `json:"teamsRefs,omitempty" tf:"-"`

	// Selector for a list of github_team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsSelector *v1.Selector `json:"teamsSelector,omitempty" tf:"-"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github_user
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*float64 `json:"users,omitempty" tf:"users,omitempty"`

	// References to github_user to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of github_user to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Environment is the Schema for the Environments API. Creates and manages environments for GitHub repositories
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
