package vpc


type VpnGatewayRoutePropagationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpn_gateway_route_propagation#create VpnGatewayRoutePropagation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpn_gateway_route_propagation#delete VpnGatewayRoutePropagation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

