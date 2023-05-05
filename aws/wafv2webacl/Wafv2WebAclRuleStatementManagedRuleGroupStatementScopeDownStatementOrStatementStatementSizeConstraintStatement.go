package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementSizeConstraintStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#comparison_operator Wafv2WebAcl#comparison_operator}.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#size Wafv2WebAcl#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#text_transformation Wafv2WebAcl#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/wafv2_web_acl#field_to_match Wafv2WebAcl#field_to_match}
	FieldToMatch *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementSizeConstraintStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

