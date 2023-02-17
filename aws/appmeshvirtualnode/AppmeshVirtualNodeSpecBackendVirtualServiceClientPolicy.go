package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

