// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2localgateways


type DataAwsEc2LocalGatewaysFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/data-sources/ec2_local_gateways#name DataAwsEc2LocalGateways#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/data-sources/ec2_local_gateways#values DataAwsEc2LocalGateways#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

