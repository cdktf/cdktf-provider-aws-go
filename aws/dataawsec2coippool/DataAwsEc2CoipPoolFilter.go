// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2coippool


type DataAwsEc2CoipPoolFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/data-sources/ec2_coip_pool#name DataAwsEc2CoipPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/data-sources/ec2_coip_pool#values DataAwsEc2CoipPool#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

