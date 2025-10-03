// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleActionCaptchaCustomRequestHandlingInsertHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/wafv2_web_acl#name Wafv2WebAcl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/wafv2_web_acl#value Wafv2WebAcl#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

