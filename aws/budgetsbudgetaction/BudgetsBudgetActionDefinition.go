// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetaction


type BudgetsBudgetActionDefinition struct {
	// iam_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.37.0/docs/resources/budgets_budget_action#iam_action_definition BudgetsBudgetAction#iam_action_definition}
	IamActionDefinition *BudgetsBudgetActionDefinitionIamActionDefinition `field:"optional" json:"iamActionDefinition" yaml:"iamActionDefinition"`
	// scp_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.37.0/docs/resources/budgets_budget_action#scp_action_definition BudgetsBudgetAction#scp_action_definition}
	ScpActionDefinition *BudgetsBudgetActionDefinitionScpActionDefinition `field:"optional" json:"scpActionDefinition" yaml:"scpActionDefinition"`
	// ssm_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.37.0/docs/resources/budgets_budget_action#ssm_action_definition BudgetsBudgetAction#ssm_action_definition}
	SsmActionDefinition *BudgetsBudgetActionDefinitionSsmActionDefinition `field:"optional" json:"ssmActionDefinition" yaml:"ssmActionDefinition"`
}

