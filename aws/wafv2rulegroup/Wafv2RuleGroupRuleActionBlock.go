// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleActionBlock struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_rule_group#custom_response Wafv2RuleGroup#custom_response}
	CustomResponse *Wafv2RuleGroupRuleActionBlockCustomResponse `field:"optional" json:"customResponse" yaml:"customResponse"`
}

