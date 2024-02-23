// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsapigatewayv2export

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsApigatewayv2ExportConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#api_id DataAwsApigatewayv2Export#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#output_type DataAwsApigatewayv2Export#output_type}.
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#specification DataAwsApigatewayv2Export#specification}.
	Specification *string `field:"required" json:"specification" yaml:"specification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#export_version DataAwsApigatewayv2Export#export_version}.
	ExportVersion *string `field:"optional" json:"exportVersion" yaml:"exportVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#id DataAwsApigatewayv2Export#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#include_extensions DataAwsApigatewayv2Export#include_extensions}.
	IncludeExtensions interface{} `field:"optional" json:"includeExtensions" yaml:"includeExtensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/data-sources/apigatewayv2_export#stage_name DataAwsApigatewayv2Export#stage_name}.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

