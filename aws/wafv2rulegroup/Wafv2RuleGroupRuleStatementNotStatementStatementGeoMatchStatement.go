package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatementStatementGeoMatchStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/wafv2_rule_group#country_codes Wafv2RuleGroup#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/wafv2_rule_group#forwarded_ip_config Wafv2RuleGroup#forwarded_ip_config}
	ForwardedIpConfig *Wafv2RuleGroupRuleStatementNotStatementStatementGeoMatchStatementForwardedIpConfig `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

