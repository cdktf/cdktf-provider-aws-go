// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetCostFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/budgets_budget#name BudgetsBudget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/budgets_budget#values BudgetsBudget#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

