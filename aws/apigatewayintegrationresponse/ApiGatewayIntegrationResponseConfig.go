// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayintegrationresponse

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayIntegrationResponseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#http_method ApiGatewayIntegrationResponse#http_method}.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#resource_id ApiGatewayIntegrationResponse#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#rest_api_id ApiGatewayIntegrationResponse#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#status_code ApiGatewayIntegrationResponse#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#content_handling ApiGatewayIntegrationResponse#content_handling}.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#id ApiGatewayIntegrationResponse#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#response_parameters ApiGatewayIntegrationResponse#response_parameters}.
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#response_templates ApiGatewayIntegrationResponse#response_templates}.
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/api_gateway_integration_response#selection_pattern ApiGatewayIntegrationResponse#selection_pattern}.
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

