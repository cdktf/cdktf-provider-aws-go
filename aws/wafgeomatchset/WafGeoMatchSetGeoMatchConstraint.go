// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafgeomatchset


type WafGeoMatchSetGeoMatchConstraint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/waf_geo_match_set#type WafGeoMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/waf_geo_match_set#value WafGeoMatchSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

