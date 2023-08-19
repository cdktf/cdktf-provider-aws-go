package apigatewayv2api

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2ApiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#name Apigatewayv2Api#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#protocol_type Apigatewayv2Api#protocol_type}.
	ProtocolType *string `field:"required" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#api_key_selection_expression Apigatewayv2Api#api_key_selection_expression}.
	ApiKeySelectionExpression *string `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#body Apigatewayv2Api#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// cors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#cors_configuration Apigatewayv2Api#cors_configuration}
	CorsConfiguration *Apigatewayv2ApiCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#credentials_arn Apigatewayv2Api#credentials_arn}.
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#description Apigatewayv2Api#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#disable_execute_api_endpoint Apigatewayv2Api#disable_execute_api_endpoint}.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#fail_on_warnings Apigatewayv2Api#fail_on_warnings}.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#id Apigatewayv2Api#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#route_key Apigatewayv2Api#route_key}.
	RouteKey *string `field:"optional" json:"routeKey" yaml:"routeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#route_selection_expression Apigatewayv2Api#route_selection_expression}.
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#tags Apigatewayv2Api#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#tags_all Apigatewayv2Api#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#target Apigatewayv2Api#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/apigatewayv2_api#version Apigatewayv2Api#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

