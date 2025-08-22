// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2authorizer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2AuthorizerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#api_id Apigatewayv2Authorizer#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#authorizer_type Apigatewayv2Authorizer#authorizer_type}.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#name Apigatewayv2Authorizer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#authorizer_credentials_arn Apigatewayv2Authorizer#authorizer_credentials_arn}.
	AuthorizerCredentialsArn *string `field:"optional" json:"authorizerCredentialsArn" yaml:"authorizerCredentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#authorizer_payload_format_version Apigatewayv2Authorizer#authorizer_payload_format_version}.
	AuthorizerPayloadFormatVersion *string `field:"optional" json:"authorizerPayloadFormatVersion" yaml:"authorizerPayloadFormatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#authorizer_result_ttl_in_seconds Apigatewayv2Authorizer#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#authorizer_uri Apigatewayv2Authorizer#authorizer_uri}.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#enable_simple_responses Apigatewayv2Authorizer#enable_simple_responses}.
	EnableSimpleResponses interface{} `field:"optional" json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#id Apigatewayv2Authorizer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#identity_sources Apigatewayv2Authorizer#identity_sources}.
	IdentitySources *[]*string `field:"optional" json:"identitySources" yaml:"identitySources"`
	// jwt_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#jwt_configuration Apigatewayv2Authorizer#jwt_configuration}
	JwtConfiguration *Apigatewayv2AuthorizerJwtConfiguration `field:"optional" json:"jwtConfiguration" yaml:"jwtConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#region Apigatewayv2Authorizer#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/apigatewayv2_authorizer#timeouts Apigatewayv2Authorizer#timeouts}
	Timeouts *Apigatewayv2AuthorizerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

