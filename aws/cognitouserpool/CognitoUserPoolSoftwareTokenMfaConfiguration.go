// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSoftwareTokenMfaConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/cognito_user_pool#enabled CognitoUserPool#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

