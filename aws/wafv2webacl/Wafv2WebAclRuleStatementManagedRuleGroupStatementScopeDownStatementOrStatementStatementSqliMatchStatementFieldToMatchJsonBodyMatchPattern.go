package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBodyMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#all Wafv2WebAcl#all}
	All *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBodyMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#included_paths Wafv2WebAcl#included_paths}.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

