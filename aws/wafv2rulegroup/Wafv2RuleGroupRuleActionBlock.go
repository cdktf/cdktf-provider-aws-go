package wafv2rulegroup


type Wafv2RuleGroupRuleActionBlock struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#custom_response Wafv2RuleGroup#custom_response}
	CustomResponse *Wafv2RuleGroupRuleActionBlockCustomResponse `field:"optional" json:"customResponse" yaml:"customResponse"`
}

