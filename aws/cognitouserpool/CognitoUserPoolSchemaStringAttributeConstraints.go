// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSchemaStringAttributeConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/cognito_user_pool#max_length CognitoUserPool#max_length}.
	MaxLength *string `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/cognito_user_pool#min_length CognitoUserPool#min_length}.
	MinLength *string `field:"optional" json:"minLength" yaml:"minLength"`
}

