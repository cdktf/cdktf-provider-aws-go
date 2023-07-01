package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

