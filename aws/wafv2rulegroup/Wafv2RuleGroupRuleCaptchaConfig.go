package wafv2rulegroup


type Wafv2RuleGroupRuleCaptchaConfig struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/wafv2_rule_group#immunity_time_property Wafv2RuleGroup#immunity_time_property}
	ImmunityTimeProperty *Wafv2RuleGroupRuleCaptchaConfigImmunityTimeProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

