package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

