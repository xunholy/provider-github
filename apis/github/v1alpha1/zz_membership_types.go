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

type MembershipInitParameters struct {

	// Defaults to false. If set to true,
	// when this resource is destroyed, the member will not be removed
	// from the organization. Instead, the member's role will be
	// downgraded to 'member'.
	// Instead of removing the member from the org, you can choose to downgrade their membership to 'member' when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization
	DowngradeOnDestroy *bool `json:"downgradeOnDestroy,omitempty" tf:"downgrade_on_destroy,omitempty"`

	// The role of the user within the organization.
	// Must be one of member or admin. Defaults to member.
	// admin role represents the owner role available via GitHub UI.
	// The role of the user within the organization. Must be one of 'member' or 'admin'.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the organization.
	// The user to add to the organization.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MembershipObservation struct {

	// Defaults to false. If set to true,
	// when this resource is destroyed, the member will not be removed
	// from the organization. Instead, the member's role will be
	// downgraded to 'member'.
	// Instead of removing the member from the org, you can choose to downgrade their membership to 'member' when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization
	DowngradeOnDestroy *bool `json:"downgradeOnDestroy,omitempty" tf:"downgrade_on_destroy,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The role of the user within the organization.
	// Must be one of member or admin. Defaults to member.
	// admin role represents the owner role available via GitHub UI.
	// The role of the user within the organization. Must be one of 'member' or 'admin'.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the organization.
	// The user to add to the organization.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MembershipParameters struct {

	// Defaults to false. If set to true,
	// when this resource is destroyed, the member will not be removed
	// from the organization. Instead, the member's role will be
	// downgraded to 'member'.
	// Instead of removing the member from the org, you can choose to downgrade their membership to 'member' when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization
	// +kubebuilder:validation:Optional
	DowngradeOnDestroy *bool `json:"downgradeOnDestroy,omitempty" tf:"downgrade_on_destroy,omitempty"`

	// The role of the user within the organization.
	// Must be one of member or admin. Defaults to member.
	// admin role represents the owner role available via GitHub UI.
	// The role of the user within the organization. Must be one of 'member' or 'admin'.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the organization.
	// The user to add to the organization.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// MembershipSpec defines the desired state of Membership
type MembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembershipParameters `json:"forProvider"`
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
	InitProvider MembershipInitParameters `json:"initProvider,omitempty"`
}

// MembershipStatus defines the observed state of Membership.
type MembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Membership is the Schema for the Memberships API. Provides a GitHub membership resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Membership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   MembershipSpec   `json:"spec"`
	Status MembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembershipList contains a list of Memberships
type MembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Membership `json:"items"`
}

// Repository type metadata.
var (
	Membership_Kind             = "Membership"
	Membership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Membership_Kind}.String()
	Membership_KindAPIVersion   = Membership_Kind + "." + CRDGroupVersion.String()
	Membership_GroupVersionKind = CRDGroupVersion.WithKind(Membership_Kind)
)

func init() {
	SchemeBuilder.Register(&Membership{}, &MembershipList{})
}