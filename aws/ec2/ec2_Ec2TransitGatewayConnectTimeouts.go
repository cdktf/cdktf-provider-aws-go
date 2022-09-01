package ec2


type Ec2TransitGatewayConnectTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_transit_gateway_connect#create Ec2TransitGatewayConnect#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_transit_gateway_connect#delete Ec2TransitGatewayConnect#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_transit_gateway_connect#update Ec2TransitGatewayConnect#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

