package appmeshvirtualrouter


type AppmeshVirtualRouterSpecListener struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_router#port_mapping AppmeshVirtualRouter#port_mapping}
	PortMapping *AppmeshVirtualRouterSpecListenerPortMapping `field:"required" json:"portMapping" yaml:"portMapping"`
}

