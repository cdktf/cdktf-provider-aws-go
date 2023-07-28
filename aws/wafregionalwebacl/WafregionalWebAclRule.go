package wafregionalwebacl


type WafregionalWebAclRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/wafregional_web_acl#priority WafregionalWebAcl#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/wafregional_web_acl#rule_id WafregionalWebAcl#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/wafregional_web_acl#action WafregionalWebAcl#action}
	Action *WafregionalWebAclRuleAction `field:"optional" json:"action" yaml:"action"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/wafregional_web_acl#override_action WafregionalWebAcl#override_action}
	OverrideAction *WafregionalWebAclRuleOverrideAction `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/wafregional_web_acl#type WafregionalWebAcl#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

