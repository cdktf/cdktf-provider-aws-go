// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apimapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2ApiMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#api_id Apigatewayv2ApiMapping#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#domain_name Apigatewayv2ApiMapping#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#stage Apigatewayv2ApiMapping#stage}.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#api_mapping_key Apigatewayv2ApiMapping#api_mapping_key}.
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#id Apigatewayv2ApiMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/apigatewayv2_api_mapping#region Apigatewayv2ApiMapping#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

