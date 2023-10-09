// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/wafv2_web_acl#arn Wafv2WebAcl#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/wafv2_web_acl#text_transformation Wafv2WebAcl#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/wafv2_web_acl#field_to_match Wafv2WebAcl#field_to_match}
	FieldToMatch *Wafv2WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

