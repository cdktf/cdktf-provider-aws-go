// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSchemaNumberAttributeConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/cognito_user_pool#max_value CognitoUserPool#max_value}.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/cognito_user_pool#min_value CognitoUserPool#min_value}.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

