package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackend struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#virtual_service AppmeshVirtualNode#virtual_service}
	VirtualService *AppmeshVirtualNodeSpecBackendVirtualService `field:"required" json:"virtualService" yaml:"virtualService"`
}

