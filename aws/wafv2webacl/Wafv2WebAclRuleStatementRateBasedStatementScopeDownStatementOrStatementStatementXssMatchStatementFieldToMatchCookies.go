// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies struct {
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafv2_web_acl#match_pattern Wafv2WebAcl#match_pattern}
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafv2_web_acl#match_scope Wafv2WebAcl#match_scope}.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafv2_web_acl#oversize_handling Wafv2WebAcl#oversize_handling}.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

