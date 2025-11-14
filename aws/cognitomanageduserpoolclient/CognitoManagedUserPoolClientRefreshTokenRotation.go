// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitomanageduserpoolclient


type CognitoManagedUserPoolClientRefreshTokenRotation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/cognito_managed_user_pool_client#feature CognitoManagedUserPoolClient#feature}.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/cognito_managed_user_pool_client#retry_grace_period_seconds CognitoManagedUserPoolClient#retry_grace_period_seconds}.
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

