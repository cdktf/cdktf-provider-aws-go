package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatement struct {
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#byte_match_statement Wafv2WebAcl#byte_match_statement}
	ByteMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementByteMatchStatement `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#geo_match_statement Wafv2WebAcl#geo_match_statement}
	GeoMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementGeoMatchStatement `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#ip_set_reference_statement Wafv2WebAcl#ip_set_reference_statement}
	IpSetReferenceStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementIpSetReferenceStatement `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#label_match_statement Wafv2WebAcl#label_match_statement}
	LabelMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementLabelMatchStatement `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#regex_match_statement Wafv2WebAcl#regex_match_statement}
	RegexMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementRegexMatchStatement `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#regex_pattern_set_reference_statement Wafv2WebAcl#regex_pattern_set_reference_statement}
	RegexPatternSetReferenceStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementRegexPatternSetReferenceStatement `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#size_constraint_statement Wafv2WebAcl#size_constraint_statement}
	SizeConstraintStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementSizeConstraintStatement `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#sqli_match_statement Wafv2WebAcl#sqli_match_statement}
	SqliMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementSqliMatchStatement `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#xss_match_statement Wafv2WebAcl#xss_match_statement}
	XssMatchStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementOrStatementStatementXssMatchStatement `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

