package wafwebacl


type WafWebAclLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_web_acl#log_destination WafWebAcl#log_destination}.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_web_acl#redacted_fields WafWebAcl#redacted_fields}
	RedactedFields *WafWebAclLoggingConfigurationRedactedFields `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

