package networkfirewall


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#action NetworkfirewallRuleGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#header NetworkfirewallRuleGroup#header}
	Header *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleHeader `field:"required" json:"header" yaml:"header"`
	// rule_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#rule_option NetworkfirewallRuleGroup#rule_option}
	RuleOption interface{} `field:"required" json:"ruleOption" yaml:"ruleOption"`
}
