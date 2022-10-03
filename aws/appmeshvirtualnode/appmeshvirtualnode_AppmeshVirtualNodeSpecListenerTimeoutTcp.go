package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTimeoutTcp struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#idle AppmeshVirtualNode#idle}
	Idle *AppmeshVirtualNodeSpecListenerTimeoutTcpIdle `field:"optional" json:"idle" yaml:"idle"`
}

