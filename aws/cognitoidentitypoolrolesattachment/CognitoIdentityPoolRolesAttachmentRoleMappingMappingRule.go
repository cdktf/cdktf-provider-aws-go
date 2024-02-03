// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypoolrolesattachment


type CognitoIdentityPoolRolesAttachmentRoleMappingMappingRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_identity_pool_roles_attachment#claim CognitoIdentityPoolRolesAttachment#claim}.
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_identity_pool_roles_attachment#match_type CognitoIdentityPoolRolesAttachment#match_type}.
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_identity_pool_roles_attachment#role_arn CognitoIdentityPoolRolesAttachment#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_identity_pool_roles_attachment#value CognitoIdentityPoolRolesAttachment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

