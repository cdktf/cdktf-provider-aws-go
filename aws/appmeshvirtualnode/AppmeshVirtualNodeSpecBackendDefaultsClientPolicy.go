package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

