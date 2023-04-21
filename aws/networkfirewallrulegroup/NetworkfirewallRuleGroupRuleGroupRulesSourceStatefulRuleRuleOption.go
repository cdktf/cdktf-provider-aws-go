package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleRuleOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/networkfirewall_rule_group#keyword NetworkfirewallRuleGroup#keyword}.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/networkfirewall_rule_group#settings NetworkfirewallRuleGroup#settings}.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

