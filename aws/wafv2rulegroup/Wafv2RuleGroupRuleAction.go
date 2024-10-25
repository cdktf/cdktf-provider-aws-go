// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_rule_group#allow Wafv2RuleGroup#allow}
	Allow *Wafv2RuleGroupRuleActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_rule_group#block Wafv2RuleGroup#block}
	Block *Wafv2RuleGroupRuleActionBlock `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_rule_group#captcha Wafv2RuleGroup#captcha}
	Captcha *Wafv2RuleGroupRuleActionCaptcha `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_rule_group#challenge Wafv2RuleGroup#challenge}
	Challenge *Wafv2RuleGroupRuleActionChallenge `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_rule_group#count Wafv2RuleGroup#count}
	Count *Wafv2RuleGroupRuleActionCount `field:"optional" json:"count" yaml:"count"`
}

