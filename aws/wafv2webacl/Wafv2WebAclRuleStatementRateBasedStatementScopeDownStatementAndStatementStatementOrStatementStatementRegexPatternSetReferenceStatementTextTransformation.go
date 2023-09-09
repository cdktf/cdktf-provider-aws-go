// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/wafv2_web_acl#priority Wafv2WebAcl#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/wafv2_web_acl#type Wafv2WebAcl#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

