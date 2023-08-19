package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTls struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_gateway#validation AppmeshVirtualGateway#validation}
	Validation *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_gateway#certificate AppmeshVirtualGateway#certificate}
	Certificate *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_gateway#enforce AppmeshVirtualGateway#enforce}.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_gateway#ports AppmeshVirtualGateway#ports}.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

