package ec2


type DataAwsEc2TransitGatewayVpcAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_vpc_attachment#name DataAwsEc2TransitGatewayVpcAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_vpc_attachment#values DataAwsEc2TransitGatewayVpcAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

