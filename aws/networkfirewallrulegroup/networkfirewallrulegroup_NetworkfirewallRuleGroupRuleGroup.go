package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroup struct {
	// rules_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#rules_source NetworkfirewallRuleGroup#rules_source}
	RulesSource *NetworkfirewallRuleGroupRuleGroupRulesSource `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// rule_variables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#rule_variables NetworkfirewallRuleGroup#rule_variables}
	RuleVariables *NetworkfirewallRuleGroupRuleGroupRuleVariables `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// stateful_rule_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#stateful_rule_options NetworkfirewallRuleGroup#stateful_rule_options}
	StatefulRuleOptions *NetworkfirewallRuleGroupRuleGroupStatefulRuleOptions `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}
