package appmesh


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#sds AppmeshVirtualNode#sds}
	Sds *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds `field:"optional" json:"sds" yaml:"sds"`
}

