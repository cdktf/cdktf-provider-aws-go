package dataawsvpngateway


type DataAwsVpnGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/data-sources/vpn_gateway#name DataAwsVpnGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/data-sources/vpn_gateway#values DataAwsVpnGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

