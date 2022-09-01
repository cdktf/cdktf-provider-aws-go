package vpc


type DataAwsVpnGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpn_gateway#name DataAwsVpnGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpn_gateway#values DataAwsVpnGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

