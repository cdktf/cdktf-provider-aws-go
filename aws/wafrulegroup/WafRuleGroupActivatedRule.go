// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafrulegroup


type WafRuleGroupActivatedRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/waf_rule_group#action WafRuleGroup#action}
	Action *WafRuleGroupActivatedRuleAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/waf_rule_group#priority WafRuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/waf_rule_group#rule_id WafRuleGroup#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/waf_rule_group#type WafRuleGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

