package ec2


type DataAwsEc2ClientVpnEndpointFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_client_vpn_endpoint#name DataAwsEc2ClientVpnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_client_vpn_endpoint#values DataAwsEc2ClientVpnEndpoint#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

