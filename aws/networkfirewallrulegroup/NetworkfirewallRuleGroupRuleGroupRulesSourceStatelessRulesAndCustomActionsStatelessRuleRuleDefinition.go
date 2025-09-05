// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/networkfirewall_rule_group#actions NetworkfirewallRuleGroup#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// match_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/networkfirewall_rule_group#match_attributes NetworkfirewallRuleGroup#match_attributes}
	MatchAttributes *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributes `field:"required" json:"matchAttributes" yaml:"matchAttributes"`
}

