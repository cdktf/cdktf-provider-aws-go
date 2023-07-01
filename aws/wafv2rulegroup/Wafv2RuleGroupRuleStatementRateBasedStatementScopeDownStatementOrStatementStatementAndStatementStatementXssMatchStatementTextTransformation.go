package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementXssMatchStatementTextTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/wafv2_rule_group#priority Wafv2RuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/wafv2_rule_group#type Wafv2RuleGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

