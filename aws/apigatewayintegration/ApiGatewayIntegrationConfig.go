// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#http_method ApiGatewayIntegration#http_method}.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#resource_id ApiGatewayIntegration#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#rest_api_id ApiGatewayIntegration#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#type ApiGatewayIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#cache_key_parameters ApiGatewayIntegration#cache_key_parameters}.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#cache_namespace ApiGatewayIntegration#cache_namespace}.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#connection_id ApiGatewayIntegration#connection_id}.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#connection_type ApiGatewayIntegration#connection_type}.
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#content_handling ApiGatewayIntegration#content_handling}.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#credentials ApiGatewayIntegration#credentials}.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#id ApiGatewayIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#integration_http_method ApiGatewayIntegration#integration_http_method}.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#passthrough_behavior ApiGatewayIntegration#passthrough_behavior}.
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#region ApiGatewayIntegration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#request_parameters ApiGatewayIntegration#request_parameters}.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#request_templates ApiGatewayIntegration#request_templates}.
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#timeout_milliseconds ApiGatewayIntegration#timeout_milliseconds}.
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#tls_config ApiGatewayIntegration#tls_config}
	TlsConfig *ApiGatewayIntegrationTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/api_gateway_integration#uri ApiGatewayIntegration#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

