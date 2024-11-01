// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection struct {
	// access_token_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/verifiedpermissions_identity_source#access_token_only VerifiedpermissionsIdentitySource#access_token_only}
	AccessTokenOnly interface{} `field:"optional" json:"accessTokenOnly" yaml:"accessTokenOnly"`
	// identity_token_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/verifiedpermissions_identity_source#identity_token_only VerifiedpermissionsIdentitySource#identity_token_only}
	IdentityTokenOnly interface{} `field:"optional" json:"identityTokenOnly" yaml:"identityTokenOnly"`
}

