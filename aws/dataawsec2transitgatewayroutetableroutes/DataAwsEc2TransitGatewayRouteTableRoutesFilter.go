// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayroutetableroutes


type DataAwsEc2TransitGatewayRouteTableRoutesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ec2_transit_gateway_route_table_routes#name DataAwsEc2TransitGatewayRouteTableRoutes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ec2_transit_gateway_route_table_routes#values DataAwsEc2TransitGatewayRouteTableRoutes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

