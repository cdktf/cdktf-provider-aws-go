// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategorySplitChargeRuleParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ce_cost_category#type CeCostCategory#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ce_cost_category#values CeCostCategory#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

