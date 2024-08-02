// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafxssmatchset


type WafXssMatchSetXssMatchTuplesFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/waf_xss_match_set#type WafXssMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/waf_xss_match_set#data WafXssMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

