// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetsBudgetActionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// action_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#action_threshold BudgetsBudgetAction#action_threshold}
	ActionThreshold *BudgetsBudgetActionActionThreshold `field:"required" json:"actionThreshold" yaml:"actionThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#action_type BudgetsBudgetAction#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#approval_model BudgetsBudgetAction#approval_model}.
	ApprovalModel *string `field:"required" json:"approvalModel" yaml:"approvalModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#budget_name BudgetsBudgetAction#budget_name}.
	BudgetName *string `field:"required" json:"budgetName" yaml:"budgetName"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#definition BudgetsBudgetAction#definition}
	Definition *BudgetsBudgetActionDefinition `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#execution_role_arn BudgetsBudgetAction#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#notification_type BudgetsBudgetAction#notification_type}.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// subscriber block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#subscriber BudgetsBudgetAction#subscriber}
	Subscriber interface{} `field:"required" json:"subscriber" yaml:"subscriber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#account_id BudgetsBudgetAction#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#id BudgetsBudgetAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#tags BudgetsBudgetAction#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#tags_all BudgetsBudgetAction#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget_action#timeouts BudgetsBudgetAction#timeouts}
	Timeouts *BudgetsBudgetActionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

