package vpc


type NetworkmanagerTransitGatewayPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_transit_gateway_peering#create NetworkmanagerTransitGatewayPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_transit_gateway_peering#delete NetworkmanagerTransitGatewayPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

