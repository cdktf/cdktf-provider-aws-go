// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSource struct {
	// rules_source_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/networkfirewall_rule_group#rules_source_list NetworkfirewallRuleGroup#rules_source_list}
	RulesSourceList *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/networkfirewall_rule_group#rules_string NetworkfirewallRuleGroup#rules_string}.
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// stateful_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/networkfirewall_rule_group#stateful_rule NetworkfirewallRuleGroup#stateful_rule}
	StatefulRule interface{} `field:"optional" json:"statefulRule" yaml:"statefulRule"`
	// stateless_rules_and_custom_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/networkfirewall_rule_group#stateless_rules_and_custom_actions NetworkfirewallRuleGroup#stateless_rules_and_custom_actions}
	StatelessRulesAndCustomActions *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}

