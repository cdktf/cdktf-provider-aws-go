package dataawsec2transitgatewayroutetableassociations


type DataAwsEc2TransitGatewayRouteTableAssociationsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/ec2_transit_gateway_route_table_associations#name DataAwsEc2TransitGatewayRouteTableAssociations#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/ec2_transit_gateway_route_table_associations#values DataAwsEc2TransitGatewayRouteTableAssociations#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

