package appmesh


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#trust AppmeshVirtualNode#trust}
	Trust *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#subject_alternative_names AppmeshVirtualNode#subject_alternative_names}
	SubjectAlternativeNames *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

