package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementRegexPatternSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_rule_group#arn Wafv2RuleGroup#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_rule_group#text_transformation Wafv2RuleGroup#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_rule_group#field_to_match Wafv2RuleGroup#field_to_match}
	FieldToMatch *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementRegexPatternSetReferenceStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

