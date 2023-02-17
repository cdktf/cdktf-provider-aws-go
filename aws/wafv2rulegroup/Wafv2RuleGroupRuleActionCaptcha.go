package wafv2rulegroup


type Wafv2RuleGroupRuleActionCaptcha struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#custom_request_handling Wafv2RuleGroup#custom_request_handling}
	CustomRequestHandling *Wafv2RuleGroupRuleActionCaptchaCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

