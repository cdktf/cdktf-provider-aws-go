package wafv2rulegroup


type Wafv2RuleGroupRuleStatementXssMatchStatement struct {
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#text_transformation Wafv2RuleGroup#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#field_to_match Wafv2RuleGroup#field_to_match}
	FieldToMatch *Wafv2RuleGroupRuleStatementXssMatchStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}
