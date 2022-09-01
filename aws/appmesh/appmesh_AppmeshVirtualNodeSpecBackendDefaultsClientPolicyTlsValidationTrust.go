package appmesh


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrust struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#acm AppmeshVirtualNode#acm}
	Acm *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcm `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#sds AppmeshVirtualNode#sds}
	Sds *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds `field:"optional" json:"sds" yaml:"sds"`
}

