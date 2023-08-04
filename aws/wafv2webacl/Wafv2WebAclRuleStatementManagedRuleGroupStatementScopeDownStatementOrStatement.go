package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/wafv2_web_acl#statement Wafv2WebAcl#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

