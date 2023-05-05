package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceList struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/networkfirewall_rule_group#generated_rules_type NetworkfirewallRuleGroup#generated_rules_type}.
	GeneratedRulesType *string `field:"required" json:"generatedRulesType" yaml:"generatedRulesType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/networkfirewall_rule_group#targets NetworkfirewallRuleGroup#targets}.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/networkfirewall_rule_group#target_types NetworkfirewallRuleGroup#target_types}.
	TargetTypes *[]*string `field:"required" json:"targetTypes" yaml:"targetTypes"`
}

