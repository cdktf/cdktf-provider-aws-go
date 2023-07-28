package wafwebacl


type WafWebAclRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/waf_web_acl#priority WafWebAcl#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/waf_web_acl#rule_id WafWebAcl#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/waf_web_acl#action WafWebAcl#action}
	Action *WafWebAclRulesAction `field:"optional" json:"action" yaml:"action"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/waf_web_acl#override_action WafWebAcl#override_action}
	OverrideAction *WafWebAclRulesOverrideAction `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/waf_web_acl#type WafWebAcl#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

