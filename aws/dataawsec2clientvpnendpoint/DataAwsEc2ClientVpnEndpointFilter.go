package dataawsec2clientvpnendpoint


type DataAwsEc2ClientVpnEndpointFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/ec2_client_vpn_endpoint#name DataAwsEc2ClientVpnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/ec2_client_vpn_endpoint#values DataAwsEc2ClientVpnEndpoint#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

