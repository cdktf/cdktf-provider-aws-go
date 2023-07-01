package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appmesh_virtual_gateway#match AppmeshVirtualGateway#match}
	Match *AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatch `field:"required" json:"match" yaml:"match"`
}

