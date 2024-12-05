// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleRuleOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/networkfirewall_rule_group#keyword NetworkfirewallRuleGroup#keyword}.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/networkfirewall_rule_group#settings NetworkfirewallRuleGroup#settings}.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

