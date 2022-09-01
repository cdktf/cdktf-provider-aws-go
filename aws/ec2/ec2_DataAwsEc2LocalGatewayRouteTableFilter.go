package ec2


type DataAwsEc2LocalGatewayRouteTableFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateway_route_table#name DataAwsEc2LocalGatewayRouteTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateway_route_table#values DataAwsEc2LocalGatewayRouteTable#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

