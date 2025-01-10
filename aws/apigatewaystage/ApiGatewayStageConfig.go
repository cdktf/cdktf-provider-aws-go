// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaystage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayStageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#deployment_id ApiGatewayStage#deployment_id}.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#rest_api_id ApiGatewayStage#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#stage_name ApiGatewayStage#stage_name}.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// access_log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#access_log_settings ApiGatewayStage#access_log_settings}
	AccessLogSettings *ApiGatewayStageAccessLogSettings `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#cache_cluster_enabled ApiGatewayStage#cache_cluster_enabled}.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#cache_cluster_size ApiGatewayStage#cache_cluster_size}.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// canary_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#canary_settings ApiGatewayStage#canary_settings}
	CanarySettings *ApiGatewayStageCanarySettings `field:"optional" json:"canarySettings" yaml:"canarySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#client_certificate_id ApiGatewayStage#client_certificate_id}.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#description ApiGatewayStage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#documentation_version ApiGatewayStage#documentation_version}.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#id ApiGatewayStage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#tags ApiGatewayStage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#tags_all ApiGatewayStage#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#variables ApiGatewayStage#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/api_gateway_stage#xray_tracing_enabled ApiGatewayStage#xray_tracing_enabled}.
	XrayTracingEnabled interface{} `field:"optional" json:"xrayTracingEnabled" yaml:"xrayTracingEnabled"`
}

