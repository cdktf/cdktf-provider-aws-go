// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/networkfirewall_rule_group#priority NetworkfirewallRuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/networkfirewall_rule_group#rule_definition NetworkfirewallRuleGroup#rule_definition}
	RuleDefinition *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinition `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
}

