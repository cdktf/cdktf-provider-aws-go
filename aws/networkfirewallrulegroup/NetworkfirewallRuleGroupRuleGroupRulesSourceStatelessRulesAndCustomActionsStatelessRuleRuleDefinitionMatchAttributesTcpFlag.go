// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesTcpFlag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/networkfirewall_rule_group#flags NetworkfirewallRuleGroup#flags}.
	Flags *[]*string `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/networkfirewall_rule_group#masks NetworkfirewallRuleGroup#masks}.
	Masks *[]*string `field:"optional" json:"masks" yaml:"masks"`
}

