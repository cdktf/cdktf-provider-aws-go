package wafv2rulegroup


type Wafv2RuleGroupRuleActionAllow struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_rule_group#custom_request_handling Wafv2RuleGroup#custom_request_handling}
	CustomRequestHandling *Wafv2RuleGroupRuleActionAllowCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

