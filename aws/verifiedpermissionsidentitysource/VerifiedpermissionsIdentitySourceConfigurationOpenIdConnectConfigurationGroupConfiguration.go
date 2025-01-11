// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationGroupConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/verifiedpermissions_identity_source#group_claim VerifiedpermissionsIdentitySource#group_claim}.
	GroupClaim *string `field:"required" json:"groupClaim" yaml:"groupClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/verifiedpermissions_identity_source#group_entity_type VerifiedpermissionsIdentitySource#group_entity_type}.
	GroupEntityType *string `field:"required" json:"groupEntityType" yaml:"groupEntityType"`
}

