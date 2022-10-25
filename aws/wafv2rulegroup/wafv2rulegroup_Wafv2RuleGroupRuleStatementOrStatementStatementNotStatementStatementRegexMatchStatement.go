package wafv2rulegroup


type Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementRegexMatchStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#regex_string Wafv2RuleGroup#regex_string}.
	RegexString *string `field:"required" json:"regexString" yaml:"regexString"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#text_transformation Wafv2RuleGroup#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#field_to_match Wafv2RuleGroup#field_to_match}
	FieldToMatch *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

