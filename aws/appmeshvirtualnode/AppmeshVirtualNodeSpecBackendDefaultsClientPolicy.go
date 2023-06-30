package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

