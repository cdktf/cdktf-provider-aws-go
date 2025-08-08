// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipampoolcidrallocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcIpamPoolCidrAllocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#ipam_pool_id VpcIpamPoolCidrAllocation#ipam_pool_id}.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#cidr VpcIpamPoolCidrAllocation#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#description VpcIpamPoolCidrAllocation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#disallowed_cidrs VpcIpamPoolCidrAllocation#disallowed_cidrs}.
	DisallowedCidrs *[]*string `field:"optional" json:"disallowedCidrs" yaml:"disallowedCidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#id VpcIpamPoolCidrAllocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#netmask_length VpcIpamPoolCidrAllocation#netmask_length}.
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/vpc_ipam_pool_cidr_allocation#region VpcIpamPoolCidrAllocation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

