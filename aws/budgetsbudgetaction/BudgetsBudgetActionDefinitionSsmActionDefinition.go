// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetaction


type BudgetsBudgetActionDefinitionSsmActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/budgets_budget_action#action_sub_type BudgetsBudgetAction#action_sub_type}.
	ActionSubType *string `field:"required" json:"actionSubType" yaml:"actionSubType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/budgets_budget_action#instance_ids BudgetsBudgetAction#instance_ids}.
	InstanceIds *[]*string `field:"required" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/budgets_budget_action#region BudgetsBudgetAction#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

