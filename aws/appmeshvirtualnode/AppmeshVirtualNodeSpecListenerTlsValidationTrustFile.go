package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsValidationTrustFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_node#certificate_chain AppmeshVirtualNode#certificate_chain}.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
}

