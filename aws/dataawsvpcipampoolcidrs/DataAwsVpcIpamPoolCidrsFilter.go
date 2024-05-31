// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcipampoolcidrs


type DataAwsVpcIpamPoolCidrsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/data-sources/vpc_ipam_pool_cidrs#name DataAwsVpcIpamPoolCidrs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/data-sources/vpc_ipam_pool_cidrs#values DataAwsVpcIpamPoolCidrs#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

