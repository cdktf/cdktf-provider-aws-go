// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2publicipv4pools


type DataAwsEc2PublicIpv4PoolsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/data-sources/ec2_public_ipv4_pools#name DataAwsEc2PublicIpv4Pools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/data-sources/ec2_public_ipv4_pools#values DataAwsEc2PublicIpv4Pools#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

