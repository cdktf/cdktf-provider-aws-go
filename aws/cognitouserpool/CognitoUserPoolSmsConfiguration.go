// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSmsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_user_pool#external_id CognitoUserPool#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_user_pool#sns_caller_arn CognitoUserPool#sns_caller_arn}.
	SnsCallerArn *string `field:"required" json:"snsCallerArn" yaml:"snsCallerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cognito_user_pool#sns_region CognitoUserPool#sns_region}.
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
}

