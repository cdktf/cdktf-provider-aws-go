// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2coippools


type DataAwsEc2CoipPoolsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/data-sources/ec2_coip_pools#name DataAwsEc2CoipPools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/data-sources/ec2_coip_pools#values DataAwsEc2CoipPools#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

