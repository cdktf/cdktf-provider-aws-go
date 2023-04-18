package wafv2webacl


type Wafv2WebAclRuleStatementOrStatementStatementSqliMatchStatementTextTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/wafv2_web_acl#priority Wafv2WebAcl#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/wafv2_web_acl#type Wafv2WebAcl#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

