package dataawsec2transitgatewayconnectpeer


type DataAwsEc2TransitGatewayConnectPeerFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_connect_peer#name DataAwsEc2TransitGatewayConnectPeer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_connect_peer#values DataAwsEc2TransitGatewayConnectPeer#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

