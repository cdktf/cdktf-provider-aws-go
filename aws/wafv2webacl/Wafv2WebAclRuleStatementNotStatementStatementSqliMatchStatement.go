package wafv2webacl


type Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement struct {
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/wafv2_web_acl#text_transformation Wafv2WebAcl#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/wafv2_web_acl#field_to_match Wafv2WebAcl#field_to_match}
	FieldToMatch *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

