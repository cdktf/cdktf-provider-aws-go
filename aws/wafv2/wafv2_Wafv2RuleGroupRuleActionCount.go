package wafv2


type Wafv2RuleGroupRuleActionCount struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#custom_request_handling Wafv2RuleGroup#custom_request_handling}
	CustomRequestHandling *Wafv2RuleGroupRuleActionCountCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

