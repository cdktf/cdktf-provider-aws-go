package vpc


type NetworkmanagerTransitGatewayConnectPeerAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_transit_gateway_connect_peer_association#create NetworkmanagerTransitGatewayConnectPeerAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_transit_gateway_connect_peer_association#delete NetworkmanagerTransitGatewayConnectPeerAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

