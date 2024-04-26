// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi


type AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appsync_graphql_api#user_pool_id AppsyncGraphqlApi#user_pool_id}.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appsync_graphql_api#app_id_client_regex AppsyncGraphqlApi#app_id_client_regex}.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appsync_graphql_api#aws_region AppsyncGraphqlApi#aws_region}.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
}

