package appmesh


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#trust AppmeshVirtualGateway#trust}
	Trust *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#subject_alternative_names AppmeshVirtualGateway#subject_alternative_names}
	SubjectAlternativeNames *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

