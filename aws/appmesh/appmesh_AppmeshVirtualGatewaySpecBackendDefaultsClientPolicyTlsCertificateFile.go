package appmesh


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#certificate_chain AppmeshVirtualGateway#certificate_chain}.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#private_key AppmeshVirtualGateway#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

