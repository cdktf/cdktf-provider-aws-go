// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions struct {
	// stateless_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/networkfirewall_rule_group#stateless_rule NetworkfirewallRuleGroup#stateless_rule}
	StatelessRule interface{} `field:"required" json:"statelessRule" yaml:"statelessRule"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/networkfirewall_rule_group#custom_action NetworkfirewallRuleGroup#custom_action}
	CustomAction interface{} `field:"optional" json:"customAction" yaml:"customAction"`
}

