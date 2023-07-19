package ec2transitgateway


type Ec2TransitGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/ec2_transit_gateway#create Ec2TransitGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/ec2_transit_gateway#delete Ec2TransitGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/ec2_transit_gateway#update Ec2TransitGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

