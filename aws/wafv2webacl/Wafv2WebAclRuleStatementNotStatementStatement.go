// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementNotStatementStatement struct {
	// and_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#and_statement Wafv2WebAcl#and_statement}
	AndStatement *Wafv2WebAclRuleStatementNotStatementStatementAndStatement `field:"optional" json:"andStatement" yaml:"andStatement"`
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#byte_match_statement Wafv2WebAcl#byte_match_statement}
	ByteMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementByteMatchStatement `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#geo_match_statement Wafv2WebAcl#geo_match_statement}
	GeoMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementGeoMatchStatement `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#ip_set_reference_statement Wafv2WebAcl#ip_set_reference_statement}
	IpSetReferenceStatement *Wafv2WebAclRuleStatementNotStatementStatementIpSetReferenceStatement `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#label_match_statement Wafv2WebAcl#label_match_statement}
	LabelMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementLabelMatchStatement `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// not_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#not_statement Wafv2WebAcl#not_statement}
	NotStatement *Wafv2WebAclRuleStatementNotStatement `field:"optional" json:"notStatement" yaml:"notStatement"`
	// or_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#or_statement Wafv2WebAcl#or_statement}
	OrStatement *Wafv2WebAclRuleStatementNotStatementStatementOrStatement `field:"optional" json:"orStatement" yaml:"orStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#regex_match_statement Wafv2WebAcl#regex_match_statement}
	RegexMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementRegexMatchStatement `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#regex_pattern_set_reference_statement Wafv2WebAcl#regex_pattern_set_reference_statement}
	RegexPatternSetReferenceStatement *Wafv2WebAclRuleStatementNotStatementStatementRegexPatternSetReferenceStatement `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#size_constraint_statement Wafv2WebAcl#size_constraint_statement}
	SizeConstraintStatement *Wafv2WebAclRuleStatementNotStatementStatementSizeConstraintStatement `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#sqli_match_statement Wafv2WebAcl#sqli_match_statement}
	SqliMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementSqliMatchStatement `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#xss_match_statement Wafv2WebAcl#xss_match_statement}
	XssMatchStatement *Wafv2WebAclRuleStatementNotStatementStatementXssMatchStatement `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

