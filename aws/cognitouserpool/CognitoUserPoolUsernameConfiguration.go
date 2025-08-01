// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolUsernameConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_user_pool#case_sensitive CognitoUserPool#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

