package appmesh


type AppmeshVirtualNodeSpecListenerTls struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#certificate AppmeshVirtualNode#certificate}
	Certificate *AppmeshVirtualNodeSpecListenerTlsCertificate `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#mode AppmeshVirtualNode#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#validation AppmeshVirtualNode#validation}
	Validation *AppmeshVirtualNodeSpecListenerTlsValidation `field:"optional" json:"validation" yaml:"validation"`
}

