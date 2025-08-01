// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafsizeconstraintset


type WafSizeConstraintSetSizeConstraintsFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/waf_size_constraint_set#type WafSizeConstraintSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/waf_size_constraint_set#data WafSizeConstraintSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

