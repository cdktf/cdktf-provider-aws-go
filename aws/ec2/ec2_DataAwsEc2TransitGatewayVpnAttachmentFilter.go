package ec2


type DataAwsEc2TransitGatewayVpnAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_vpn_attachment#name DataAwsEc2TransitGatewayVpnAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_vpn_attachment#values DataAwsEc2TransitGatewayVpnAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

