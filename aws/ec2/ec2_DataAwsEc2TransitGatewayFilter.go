package ec2


type DataAwsEc2TransitGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway#name DataAwsEc2TransitGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway#values DataAwsEc2TransitGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

