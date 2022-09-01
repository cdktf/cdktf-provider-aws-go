package wafv2


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementLabelMatchStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#key Wafv2WebAcl#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#scope Wafv2WebAcl#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

