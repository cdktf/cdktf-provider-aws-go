package dataawsec2transitgatewayconnect


type DataAwsEc2TransitGatewayConnectFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_transit_gateway_connect#name DataAwsEc2TransitGatewayConnect#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_transit_gateway_connect#values DataAwsEc2TransitGatewayConnect#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

