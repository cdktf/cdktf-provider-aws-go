package wafv2webacl


type Wafv2WebAclRuleOverrideAction struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#count Wafv2WebAcl#count}
	Count *Wafv2WebAclRuleOverrideActionCount `field:"optional" json:"count" yaml:"count"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#none Wafv2WebAcl#none}
	None *Wafv2WebAclRuleOverrideActionNone `field:"optional" json:"none" yaml:"none"`
}

