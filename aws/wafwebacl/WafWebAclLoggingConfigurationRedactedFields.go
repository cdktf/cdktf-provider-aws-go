package wafwebacl


type WafWebAclLoggingConfigurationRedactedFields struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/waf_web_acl#field_to_match WafWebAcl#field_to_match}
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
}

