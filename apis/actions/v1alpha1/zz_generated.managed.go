// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this EnvironmentSecret.
func (mg *EnvironmentSecret) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this EnvironmentSecret.
func (mg *EnvironmentSecret) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this EnvironmentVariable.
func (mg *EnvironmentVariable) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this EnvironmentVariable.
func (mg *EnvironmentVariable) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this OrganizationPermissions.
func (mg *OrganizationPermissions) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this OrganizationPermissions.
func (mg *OrganizationPermissions) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this OrganizationVariable.
func (mg *OrganizationVariable) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this OrganizationVariable.
func (mg *OrganizationVariable) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this OrganizationVariable.
func (mg *OrganizationVariable) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this OrganizationVariable.
func (mg *OrganizationVariable) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this OrganizationVariable.
func (mg *OrganizationVariable) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this OrganizationVariable.
func (mg *OrganizationVariable) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this OrganizationVariable.
func (mg *OrganizationVariable) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this OrganizationVariable.
func (mg *OrganizationVariable) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this OrganizationVariable.
func (mg *OrganizationVariable) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this OrganizationVariable.
func (mg *OrganizationVariable) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this OrganizationVariable.
func (mg *OrganizationVariable) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this OrganizationVariable.
func (mg *OrganizationVariable) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this RepositoryAccessLevel.
func (mg *RepositoryAccessLevel) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this RepositoryPermissions.
func (mg *RepositoryPermissions) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this RepositoryPermissions.
func (mg *RepositoryPermissions) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Secret.
func (mg *Secret) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Secret.
func (mg *Secret) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Secret.
func (mg *Secret) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Secret.
func (mg *Secret) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Secret.
func (mg *Secret) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Secret.
func (mg *Secret) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Secret.
func (mg *Secret) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Secret.
func (mg *Secret) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Secret.
func (mg *Secret) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Secret.
func (mg *Secret) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Secret.
func (mg *Secret) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Secret.
func (mg *Secret) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Variable.
func (mg *Variable) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Variable.
func (mg *Variable) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Variable.
func (mg *Variable) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Variable.
func (mg *Variable) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Variable.
func (mg *Variable) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Variable.
func (mg *Variable) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Variable.
func (mg *Variable) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Variable.
func (mg *Variable) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Variable.
func (mg *Variable) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Variable.
func (mg *Variable) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Variable.
func (mg *Variable) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Variable.
func (mg *Variable) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}