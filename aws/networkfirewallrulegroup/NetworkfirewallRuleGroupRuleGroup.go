// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroup struct {
	// rules_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/networkfirewall_rule_group#rules_source NetworkfirewallRuleGroup#rules_source}
	RulesSource *NetworkfirewallRuleGroupRuleGroupRulesSource `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// reference_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/networkfirewall_rule_group#reference_sets NetworkfirewallRuleGroup#reference_sets}
	ReferenceSets *NetworkfirewallRuleGroupRuleGroupReferenceSets `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// rule_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/networkfirewall_rule_group#rule_variables NetworkfirewallRuleGroup#rule_variables}
	RuleVariables *NetworkfirewallRuleGroupRuleGroupRuleVariables `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// stateful_rule_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/networkfirewall_rule_group#stateful_rule_options NetworkfirewallRuleGroup#stateful_rule_options}
	StatefulRuleOptions *NetworkfirewallRuleGroupRuleGroupStatefulRuleOptions `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

