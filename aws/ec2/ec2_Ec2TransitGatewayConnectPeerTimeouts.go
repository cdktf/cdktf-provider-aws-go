package ec2


type Ec2TransitGatewayConnectPeerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_transit_gateway_connect_peer#create Ec2TransitGatewayConnectPeer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_transit_gateway_connect_peer#delete Ec2TransitGatewayConnectPeer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

