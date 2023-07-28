package dataawsec2transitgatewaypeeringattachment


type DataAwsEc2TransitGatewayPeeringAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/ec2_transit_gateway_peering_attachment#name DataAwsEc2TransitGatewayPeeringAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/ec2_transit_gateway_peering_attachment#values DataAwsEc2TransitGatewayPeeringAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

