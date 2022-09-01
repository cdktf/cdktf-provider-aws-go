package ec2


type Ec2ClientVpnEndpointAuthenticationOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_endpoint#type Ec2ClientVpnEndpoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_endpoint#active_directory_id Ec2ClientVpnEndpoint#active_directory_id}.
	ActiveDirectoryId *string `field:"optional" json:"activeDirectoryId" yaml:"activeDirectoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_endpoint#root_certificate_chain_arn Ec2ClientVpnEndpoint#root_certificate_chain_arn}.
	RootCertificateChainArn *string `field:"optional" json:"rootCertificateChainArn" yaml:"rootCertificateChainArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_endpoint#saml_provider_arn Ec2ClientVpnEndpoint#saml_provider_arn}.
	SamlProviderArn *string `field:"optional" json:"samlProviderArn" yaml:"samlProviderArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_endpoint#self_service_saml_provider_arn Ec2ClientVpnEndpoint#self_service_saml_provider_arn}.
	SelfServiceSamlProviderArn *string `field:"optional" json:"selfServiceSamlProviderArn" yaml:"selfServiceSamlProviderArn"`
}

