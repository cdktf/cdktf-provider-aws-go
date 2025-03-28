// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#action Wafv2RuleGroup#action}
	Action *Wafv2RuleGroupRuleAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#name Wafv2RuleGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#priority Wafv2RuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#visibility_config Wafv2RuleGroup#visibility_config}
	VisibilityConfig *Wafv2RuleGroupRuleVisibilityConfig `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#captcha_config Wafv2RuleGroup#captcha_config}
	CaptchaConfig *Wafv2RuleGroupRuleCaptchaConfig `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#rule_label Wafv2RuleGroup#rule_label}
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/wafv2_rule_group#statement Wafv2RuleGroup#statement}
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

