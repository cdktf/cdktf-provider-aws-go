package appmesh


type AppmeshVirtualGatewaySpecListenerTlsValidationTrust struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecListenerTlsValidationTrustFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#sds AppmeshVirtualGateway#sds}
	Sds *AppmeshVirtualGatewaySpecListenerTlsValidationTrustSds `field:"optional" json:"sds" yaml:"sds"`
}

