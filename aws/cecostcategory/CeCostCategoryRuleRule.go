// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategoryRuleRule struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#and CeCostCategory#and}
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#not CeCostCategory#not}
	Not *CeCostCategoryRuleRuleNot `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#or CeCostCategory#or}
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleTags `field:"optional" json:"tags" yaml:"tags"`
}

