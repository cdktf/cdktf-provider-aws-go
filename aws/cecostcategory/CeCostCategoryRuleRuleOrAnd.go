// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategoryRuleRuleOrAnd struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleOrAndCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleOrAndDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleOrAndTags `field:"optional" json:"tags" yaml:"tags"`
}

