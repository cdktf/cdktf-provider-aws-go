package appmesh


type AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNames struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#match AppmeshVirtualNode#match}
	Match *AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch `field:"required" json:"match" yaml:"match"`
}

