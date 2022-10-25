package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatementStatementNotStatementStatementByteMatchStatementTextTransformation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#priority Wafv2RuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#type Wafv2RuleGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

