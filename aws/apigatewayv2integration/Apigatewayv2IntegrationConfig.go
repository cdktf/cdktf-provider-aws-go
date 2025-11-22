// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2integration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2IntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#api_id Apigatewayv2Integration#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#integration_type Apigatewayv2Integration#integration_type}.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#connection_id Apigatewayv2Integration#connection_id}.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#connection_type Apigatewayv2Integration#connection_type}.
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#content_handling_strategy Apigatewayv2Integration#content_handling_strategy}.
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#credentials_arn Apigatewayv2Integration#credentials_arn}.
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#description Apigatewayv2Integration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#id Apigatewayv2Integration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#integration_method Apigatewayv2Integration#integration_method}.
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#integration_subtype Apigatewayv2Integration#integration_subtype}.
	IntegrationSubtype *string `field:"optional" json:"integrationSubtype" yaml:"integrationSubtype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#integration_uri Apigatewayv2Integration#integration_uri}.
	IntegrationUri *string `field:"optional" json:"integrationUri" yaml:"integrationUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#passthrough_behavior Apigatewayv2Integration#passthrough_behavior}.
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#payload_format_version Apigatewayv2Integration#payload_format_version}.
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#region Apigatewayv2Integration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#request_parameters Apigatewayv2Integration#request_parameters}.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#request_templates Apigatewayv2Integration#request_templates}.
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// response_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#response_parameters Apigatewayv2Integration#response_parameters}
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#template_selection_expression Apigatewayv2Integration#template_selection_expression}.
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#timeout_milliseconds Apigatewayv2Integration#timeout_milliseconds}.
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apigatewayv2_integration#tls_config Apigatewayv2Integration#tls_config}
	TlsConfig *Apigatewayv2IntegrationTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

