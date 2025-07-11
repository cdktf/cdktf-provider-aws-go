// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipampoolcidr


type VpcIpamPoolCidrTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/vpc_ipam_pool_cidr#create VpcIpamPoolCidr#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/vpc_ipam_pool_cidr#delete VpcIpamPoolCidr#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

