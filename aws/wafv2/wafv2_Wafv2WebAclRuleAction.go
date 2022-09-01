package wafv2


type Wafv2WebAclRuleAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclRuleActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclRuleActionBlock `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#captcha Wafv2WebAcl#captcha}
	Captcha *Wafv2WebAclRuleActionCaptcha `field:"optional" json:"captcha" yaml:"captcha"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#count Wafv2WebAcl#count}
	Count *Wafv2WebAclRuleActionCount `field:"optional" json:"count" yaml:"count"`
}

