// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementNotStatementStatementOrStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/wafv2_web_acl#statement Wafv2WebAcl#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

