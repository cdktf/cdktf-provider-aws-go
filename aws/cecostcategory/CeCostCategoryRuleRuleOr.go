// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategoryRuleRuleOr struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleOrCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleOrDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleOrTags `field:"optional" json:"tags" yaml:"tags"`
}

