// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/api_gateway_account#cloudwatch_role_arn ApiGatewayAccount#cloudwatch_role_arn}.
	CloudwatchRoleArn *string `field:"optional" json:"cloudwatchRoleArn" yaml:"cloudwatchRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/api_gateway_account#reset_on_delete ApiGatewayAccount#reset_on_delete}.
	ResetOnDelete interface{} `field:"optional" json:"resetOnDelete" yaml:"resetOnDelete"`
}

