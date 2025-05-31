// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/verifiedpermissions_identity_source#user_pool_arn VerifiedpermissionsIdentitySource#user_pool_arn}.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/verifiedpermissions_identity_source#client_ids VerifiedpermissionsIdentitySource#client_ids}.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/verifiedpermissions_identity_source#group_configuration VerifiedpermissionsIdentitySource#group_configuration}
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
}

