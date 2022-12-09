package auditmanagercontrol


type AuditmanagerControlControlMappingSourcesSourceKeyword struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_control#keyword_input_type AuditmanagerControl#keyword_input_type}.
	KeywordInputType *string `field:"required" json:"keywordInputType" yaml:"keywordInputType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_control#keyword_value AuditmanagerControl#keyword_value}.
	KeywordValue *string `field:"required" json:"keywordValue" yaml:"keywordValue"`
}

