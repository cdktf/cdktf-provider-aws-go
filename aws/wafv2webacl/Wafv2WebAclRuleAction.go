// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclRuleActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclRuleActionBlock `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/wafv2_web_acl#captcha Wafv2WebAcl#captcha}
	Captcha *Wafv2WebAclRuleActionCaptcha `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/wafv2_web_acl#challenge Wafv2WebAcl#challenge}
	Challenge *Wafv2WebAclRuleActionChallenge `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/wafv2_web_acl#count Wafv2WebAcl#count}
	Count *Wafv2WebAclRuleActionCount `field:"optional" json:"count" yaml:"count"`
}

