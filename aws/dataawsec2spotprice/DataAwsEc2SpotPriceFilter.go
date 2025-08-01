// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2spotprice


type DataAwsEc2SpotPriceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/ec2_spot_price#name DataAwsEc2SpotPrice#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/ec2_spot_price#values DataAwsEc2SpotPrice#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

