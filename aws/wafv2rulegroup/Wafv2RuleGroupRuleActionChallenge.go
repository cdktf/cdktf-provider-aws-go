// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleActionChallenge struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/wafv2_rule_group#custom_request_handling Wafv2RuleGroup#custom_request_handling}
	CustomRequestHandling *Wafv2RuleGroupRuleActionChallengeCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

