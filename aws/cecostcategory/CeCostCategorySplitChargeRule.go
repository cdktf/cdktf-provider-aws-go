// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategorySplitChargeRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/ce_cost_category#method CeCostCategory#method}.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/ce_cost_category#source CeCostCategory#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/ce_cost_category#targets CeCostCategory#targets}.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/ce_cost_category#parameter CeCostCategory#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

