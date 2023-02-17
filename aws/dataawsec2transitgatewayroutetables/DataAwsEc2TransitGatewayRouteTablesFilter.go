package dataawsec2transitgatewayroutetables


type DataAwsEc2TransitGatewayRouteTablesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_route_tables#name DataAwsEc2TransitGatewayRouteTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_route_tables#values DataAwsEc2TransitGatewayRouteTables#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

