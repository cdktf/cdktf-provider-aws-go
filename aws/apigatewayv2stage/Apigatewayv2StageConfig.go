// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2stage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2StageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#api_id Apigatewayv2Stage#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#name Apigatewayv2Stage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#access_log_settings Apigatewayv2Stage#access_log_settings}
	AccessLogSettings *Apigatewayv2StageAccessLogSettings `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#auto_deploy Apigatewayv2Stage#auto_deploy}.
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#client_certificate_id Apigatewayv2Stage#client_certificate_id}.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// default_route_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#default_route_settings Apigatewayv2Stage#default_route_settings}
	DefaultRouteSettings *Apigatewayv2StageDefaultRouteSettings `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#deployment_id Apigatewayv2Stage#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#description Apigatewayv2Stage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#id Apigatewayv2Stage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// route_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#route_settings Apigatewayv2Stage#route_settings}
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#stage_variables Apigatewayv2Stage#stage_variables}.
	StageVariables *map[string]*string `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#tags Apigatewayv2Stage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/apigatewayv2_stage#tags_all Apigatewayv2Stage#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

