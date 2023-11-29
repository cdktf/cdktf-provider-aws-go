// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2route

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2RouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#api_id Apigatewayv2Route#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#route_key Apigatewayv2Route#route_key}.
	RouteKey *string `field:"required" json:"routeKey" yaml:"routeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#api_key_required Apigatewayv2Route#api_key_required}.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#authorization_scopes Apigatewayv2Route#authorization_scopes}.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#authorization_type Apigatewayv2Route#authorization_type}.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#authorizer_id Apigatewayv2Route#authorizer_id}.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#id Apigatewayv2Route#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#model_selection_expression Apigatewayv2Route#model_selection_expression}.
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#operation_name Apigatewayv2Route#operation_name}.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#request_models Apigatewayv2Route#request_models}.
	RequestModels *map[string]*string `field:"optional" json:"requestModels" yaml:"requestModels"`
	// request_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#request_parameter Apigatewayv2Route#request_parameter}
	RequestParameter interface{} `field:"optional" json:"requestParameter" yaml:"requestParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#route_response_selection_expression Apigatewayv2Route#route_response_selection_expression}.
	RouteResponseSelectionExpression *string `field:"optional" json:"routeResponseSelectionExpression" yaml:"routeResponseSelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/apigatewayv2_route#target Apigatewayv2Route#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

