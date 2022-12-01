package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride struct {
	// action_to_use block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#action_to_use Wafv2WebAcl#action_to_use}
	ActionToUse *Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse `field:"required" json:"actionToUse" yaml:"actionToUse"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#name Wafv2WebAcl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

