package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaults struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#client_policy AppmeshVirtualNode#client_policy}
	ClientPolicy *AppmeshVirtualNodeSpecBackendDefaultsClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

