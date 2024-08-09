// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayroutetables


type DataAwsEc2TransitGatewayRouteTablesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/data-sources/ec2_transit_gateway_route_tables#name DataAwsEc2TransitGatewayRouteTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/data-sources/ec2_transit_gateway_route_tables#values DataAwsEc2TransitGatewayRouteTables#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

