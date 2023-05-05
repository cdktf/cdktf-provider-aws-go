package dataawsec2transitgatewayroutetablepropagations


type DataAwsEc2TransitGatewayRouteTablePropagationsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_transit_gateway_route_table_propagations#name DataAwsEc2TransitGatewayRouteTablePropagations#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_transit_gateway_route_table_propagations#values DataAwsEc2TransitGatewayRouteTablePropagations#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

