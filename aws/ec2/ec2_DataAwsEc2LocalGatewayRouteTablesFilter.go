package ec2


type DataAwsEc2LocalGatewayRouteTablesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateway_route_tables#name DataAwsEc2LocalGatewayRouteTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateway_route_tables#values DataAwsEc2LocalGatewayRouteTables#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

