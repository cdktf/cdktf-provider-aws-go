package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaults struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_gateway#client_policy AppmeshVirtualGateway#client_policy}
	ClientPolicy *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

