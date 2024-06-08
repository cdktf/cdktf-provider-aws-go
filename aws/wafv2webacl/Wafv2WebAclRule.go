// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#name Wafv2WebAcl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#priority Wafv2WebAcl#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#visibility_config Wafv2WebAcl#visibility_config}
	VisibilityConfig *Wafv2WebAclRuleVisibilityConfig `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#action Wafv2WebAcl#action}
	Action *Wafv2WebAclRuleAction `field:"optional" json:"action" yaml:"action"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#captcha_config Wafv2WebAcl#captcha_config}
	CaptchaConfig *Wafv2WebAclRuleCaptchaConfig `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#override_action Wafv2WebAcl#override_action}
	OverrideAction *Wafv2WebAclRuleOverrideAction `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#rule_label Wafv2WebAcl#rule_label}
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/wafv2_web_acl#statement Wafv2WebAcl#statement}
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

