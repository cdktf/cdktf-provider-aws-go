// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcblockpublicaccessexclusion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcBlockPublicAccessExclusionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpc_block_public_access_exclusion#internet_gateway_exclusion_mode VpcBlockPublicAccessExclusion#internet_gateway_exclusion_mode}.
	InternetGatewayExclusionMode *string `field:"required" json:"internetGatewayExclusionMode" yaml:"internetGatewayExclusionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpc_block_public_access_exclusion#subnet_id VpcBlockPublicAccessExclusion#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpc_block_public_access_exclusion#tags VpcBlockPublicAccessExclusion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpc_block_public_access_exclusion#timeouts VpcBlockPublicAccessExclusion#timeouts}
	Timeouts *VpcBlockPublicAccessExclusionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpc_block_public_access_exclusion#vpc_id VpcBlockPublicAccessExclusion#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

