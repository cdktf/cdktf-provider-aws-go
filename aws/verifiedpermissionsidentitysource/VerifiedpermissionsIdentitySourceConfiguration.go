// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfiguration struct {
	// cognito_user_pool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/verifiedpermissions_identity_source#cognito_user_pool_configuration VerifiedpermissionsIdentitySource#cognito_user_pool_configuration}
	CognitoUserPoolConfiguration interface{} `field:"optional" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
	// open_id_connect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/verifiedpermissions_identity_source#open_id_connect_configuration VerifiedpermissionsIdentitySource#open_id_connect_configuration}
	OpenIdConnectConfiguration interface{} `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
}

