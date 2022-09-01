package wafregional


type WafregionalWebAclLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_web_acl#log_destination WafregionalWebAcl#log_destination}.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_web_acl#redacted_fields WafregionalWebAcl#redacted_fields}
	RedactedFields *WafregionalWebAclLoggingConfigurationRedactedFields `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

