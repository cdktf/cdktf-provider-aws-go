// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#comparison_operator BudgetsBudget#comparison_operator}.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#notification_type BudgetsBudget#notification_type}.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#threshold BudgetsBudget#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#threshold_type BudgetsBudget#threshold_type}.
	ThresholdType *string `field:"required" json:"thresholdType" yaml:"thresholdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#subscriber_email_addresses BudgetsBudget#subscriber_email_addresses}.
	SubscriberEmailAddresses *[]*string `field:"optional" json:"subscriberEmailAddresses" yaml:"subscriberEmailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/budgets_budget#subscriber_sns_topic_arns BudgetsBudget#subscriber_sns_topic_arns}.
	SubscriberSnsTopicArns *[]*string `field:"optional" json:"subscriberSnsTopicArns" yaml:"subscriberSnsTopicArns"`
}

