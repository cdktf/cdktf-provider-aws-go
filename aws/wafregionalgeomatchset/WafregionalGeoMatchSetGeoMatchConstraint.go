// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalgeomatchset


type WafregionalGeoMatchSetGeoMatchConstraint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafregional_geo_match_set#type WafregionalGeoMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafregional_geo_match_set#value WafregionalGeoMatchSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

