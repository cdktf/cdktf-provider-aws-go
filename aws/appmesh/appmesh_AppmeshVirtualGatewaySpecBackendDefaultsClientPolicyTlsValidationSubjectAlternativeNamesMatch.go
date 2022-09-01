package appmesh


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#exact AppmeshVirtualGateway#exact}.
	Exact *[]*string `field:"required" json:"exact" yaml:"exact"`
}

