package wafv2


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#statement Wafv2WebAcl#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}
