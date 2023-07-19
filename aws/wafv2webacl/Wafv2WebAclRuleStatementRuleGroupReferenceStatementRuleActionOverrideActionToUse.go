package wafv2webacl


type Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#captcha Wafv2WebAcl#captcha}
	Captcha *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha `field:"optional" json:"captcha" yaml:"captcha"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#count Wafv2WebAcl#count}
	Count *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount `field:"optional" json:"count" yaml:"count"`
}

