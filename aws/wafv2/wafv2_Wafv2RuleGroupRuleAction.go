package wafv2


type Wafv2RuleGroupRuleAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#allow Wafv2RuleGroup#allow}
	Allow *Wafv2RuleGroupRuleActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#block Wafv2RuleGroup#block}
	Block *Wafv2RuleGroupRuleActionBlock `field:"optional" json:"block" yaml:"block"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#count Wafv2RuleGroup#count}
	Count *Wafv2RuleGroupRuleActionCount `field:"optional" json:"count" yaml:"count"`
}
