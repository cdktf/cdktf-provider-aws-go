package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNames struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_node#match AppmeshVirtualNode#match}
	Match *AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch `field:"required" json:"match" yaml:"match"`
}

