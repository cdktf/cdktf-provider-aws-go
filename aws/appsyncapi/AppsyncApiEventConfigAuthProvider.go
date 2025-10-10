// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfigAuthProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_api#auth_type AppsyncApi#auth_type}.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// cognito_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_api#cognito_config AppsyncApi#cognito_config}
	CognitoConfig interface{} `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_api#lambda_authorizer_config AppsyncApi#lambda_authorizer_config}
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_api#openid_connect_config AppsyncApi#openid_connect_config}
	OpenidConnectConfig interface{} `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
}

