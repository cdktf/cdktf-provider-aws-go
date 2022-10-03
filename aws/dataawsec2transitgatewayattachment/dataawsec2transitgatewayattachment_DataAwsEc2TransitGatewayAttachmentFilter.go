package dataawsec2transitgatewayattachment


type DataAwsEc2TransitGatewayAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_attachment#name DataAwsEc2TransitGatewayAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_transit_gateway_attachment#values DataAwsEc2TransitGatewayAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

