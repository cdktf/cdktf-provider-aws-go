// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsCustomAction struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/networkfirewall_rule_group#action_definition NetworkfirewallRuleGroup#action_definition}
	ActionDefinition *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsCustomActionActionDefinition `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/networkfirewall_rule_group#action_name NetworkfirewallRuleGroup#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}

