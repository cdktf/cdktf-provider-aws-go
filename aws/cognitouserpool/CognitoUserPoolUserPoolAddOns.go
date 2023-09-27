// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolUserPoolAddOns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/cognito_user_pool#advanced_security_mode CognitoUserPool#advanced_security_mode}.
	AdvancedSecurityMode *string `field:"required" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

