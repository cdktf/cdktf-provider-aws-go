package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#tls AppmeshVirtualGateway#tls}
	Tls *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

