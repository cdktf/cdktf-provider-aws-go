// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayauthorizer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayAuthorizerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#name ApiGatewayAuthorizer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#rest_api_id ApiGatewayAuthorizer#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#authorizer_credentials ApiGatewayAuthorizer#authorizer_credentials}.
	AuthorizerCredentials *string `field:"optional" json:"authorizerCredentials" yaml:"authorizerCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#authorizer_result_ttl_in_seconds ApiGatewayAuthorizer#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#authorizer_uri ApiGatewayAuthorizer#authorizer_uri}.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#id ApiGatewayAuthorizer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#identity_source ApiGatewayAuthorizer#identity_source}.
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#identity_validation_expression ApiGatewayAuthorizer#identity_validation_expression}.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#provider_arns ApiGatewayAuthorizer#provider_arns}.
	ProviderArns *[]*string `field:"optional" json:"providerArns" yaml:"providerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/api_gateway_authorizer#type ApiGatewayAuthorizer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

