// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgateway


type DataAwsEc2TransitGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ec2_transit_gateway#name DataAwsEc2TransitGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ec2_transit_gateway#values DataAwsEc2TransitGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

