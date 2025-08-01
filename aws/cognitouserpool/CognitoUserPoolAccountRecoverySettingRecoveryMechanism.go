// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolAccountRecoverySettingRecoveryMechanism struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_user_pool#name CognitoUserPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_user_pool#priority CognitoUserPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

