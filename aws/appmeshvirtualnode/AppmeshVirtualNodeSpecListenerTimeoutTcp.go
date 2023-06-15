package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTimeoutTcp struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/appmesh_virtual_node#idle AppmeshVirtualNode#idle}
	Idle *AppmeshVirtualNodeSpecListenerTimeoutTcpIdle `field:"optional" json:"idle" yaml:"idle"`
}

