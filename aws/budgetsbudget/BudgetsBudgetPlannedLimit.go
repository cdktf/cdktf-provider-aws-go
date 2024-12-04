// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetPlannedLimit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/budgets_budget#amount BudgetsBudget#amount}.
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/budgets_budget#start_time BudgetsBudget#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/budgets_budget#unit BudgetsBudget#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

