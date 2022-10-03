package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaults struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#client_policy AppmeshVirtualGateway#client_policy}
	ClientPolicy *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

