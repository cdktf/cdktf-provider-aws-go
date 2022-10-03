package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTls struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#certificate AppmeshVirtualGateway#certificate}
	Certificate *AppmeshVirtualGatewaySpecListenerTlsCertificate `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#mode AppmeshVirtualGateway#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#validation AppmeshVirtualGateway#validation}
	Validation *AppmeshVirtualGatewaySpecListenerTlsValidation `field:"optional" json:"validation" yaml:"validation"`
}

