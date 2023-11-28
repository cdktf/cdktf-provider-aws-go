// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetAutoAdjustDataHistoricalOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/budgets_budget#budget_adjustment_period BudgetsBudget#budget_adjustment_period}.
	BudgetAdjustmentPeriod *float64 `field:"required" json:"budgetAdjustmentPeriod" yaml:"budgetAdjustmentPeriod"`
}

