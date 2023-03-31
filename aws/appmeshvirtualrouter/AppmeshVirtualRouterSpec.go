package appmeshvirtualrouter


type AppmeshVirtualRouterSpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_router#listener AppmeshVirtualRouter#listener}
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
}

