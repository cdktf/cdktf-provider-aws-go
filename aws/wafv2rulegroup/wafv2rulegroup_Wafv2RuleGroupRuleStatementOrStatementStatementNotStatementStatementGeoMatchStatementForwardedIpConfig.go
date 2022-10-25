package wafv2rulegroup


type Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementGeoMatchStatementForwardedIpConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#fallback_behavior Wafv2RuleGroup#fallback_behavior}.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#header_name Wafv2RuleGroup#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

