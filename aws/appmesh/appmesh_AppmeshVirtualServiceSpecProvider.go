package appmesh


type AppmeshVirtualServiceSpecProvider struct {
	// virtual_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_service#virtual_node AppmeshVirtualService#virtual_node}
	VirtualNode *AppmeshVirtualServiceSpecProviderVirtualNode `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// virtual_router block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_service#virtual_router AppmeshVirtualService#virtual_router}
	VirtualRouter *AppmeshVirtualServiceSpecProviderVirtualRouter `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

