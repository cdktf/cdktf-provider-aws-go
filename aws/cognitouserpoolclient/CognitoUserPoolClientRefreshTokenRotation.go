// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolclient


type CognitoUserPoolClientRefreshTokenRotation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/cognito_user_pool_client#feature CognitoUserPoolClient#feature}.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/cognito_user_pool_client#retry_grace_period_seconds CognitoUserPoolClient#retry_grace_period_seconds}.
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

