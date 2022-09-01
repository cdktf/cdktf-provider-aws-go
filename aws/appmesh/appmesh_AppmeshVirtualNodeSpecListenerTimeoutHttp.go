package appmesh


type AppmeshVirtualNodeSpecListenerTimeoutHttp struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#idle AppmeshVirtualNode#idle}
	Idle *AppmeshVirtualNodeSpecListenerTimeoutHttpIdle `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#per_request AppmeshVirtualNode#per_request}
	PerRequest *AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequest `field:"optional" json:"perRequest" yaml:"perRequest"`
}

