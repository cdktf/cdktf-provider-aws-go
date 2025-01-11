// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayroutetable


type DataAwsEc2TransitGatewayRouteTableFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/ec2_transit_gateway_route_table#name DataAwsEc2TransitGatewayRouteTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/ec2_transit_gateway_route_table#values DataAwsEc2TransitGatewayRouteTable#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

