package waf


type WafWebAclLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_web_acl#type WafWebAcl#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_web_acl#data WafWebAcl#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

