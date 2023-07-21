package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appmesh_virtual_gateway#certificate_authority_arns AppmeshVirtualGateway#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

