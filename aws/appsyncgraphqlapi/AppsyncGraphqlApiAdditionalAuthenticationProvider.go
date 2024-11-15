// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi


type AppsyncGraphqlApiAdditionalAuthenticationProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_graphql_api#authentication_type AppsyncGraphqlApi#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_graphql_api#lambda_authorizer_config AppsyncGraphqlApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_graphql_api#openid_connect_config AppsyncGraphqlApi#openid_connect_config}
	OpenidConnectConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_graphql_api#user_pool_config AppsyncGraphqlApi#user_pool_config}
	UserPoolConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

