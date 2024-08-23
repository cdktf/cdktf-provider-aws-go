// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawspricingproduct


type DataAwsPricingProductFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/data-sources/pricing_product#field DataAwsPricingProduct#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/data-sources/pricing_product#value DataAwsPricingProduct#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

