package appmesh


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

