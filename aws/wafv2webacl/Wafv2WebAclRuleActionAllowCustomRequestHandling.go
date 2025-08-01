// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleActionAllowCustomRequestHandling struct {
	// insert_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/wafv2_web_acl#insert_header Wafv2WebAcl#insert_header}
	InsertHeader interface{} `field:"required" json:"insertHeader" yaml:"insertHeader"`
}

