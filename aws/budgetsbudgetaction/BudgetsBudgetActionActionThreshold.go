// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetaction


type BudgetsBudgetActionActionThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/budgets_budget_action#action_threshold_type BudgetsBudgetAction#action_threshold_type}.
	ActionThresholdType *string `field:"required" json:"actionThresholdType" yaml:"actionThresholdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/budgets_budget_action#action_threshold_value BudgetsBudgetAction#action_threshold_value}.
	ActionThresholdValue *float64 `field:"required" json:"actionThresholdValue" yaml:"actionThresholdValue"`
}

