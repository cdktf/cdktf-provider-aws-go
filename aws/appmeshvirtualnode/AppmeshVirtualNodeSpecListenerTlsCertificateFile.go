package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsCertificateFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#certificate_chain AppmeshVirtualNode#certificate_chain}.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#private_key AppmeshVirtualNode#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

