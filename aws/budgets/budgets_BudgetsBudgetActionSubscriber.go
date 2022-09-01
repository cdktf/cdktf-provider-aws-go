package budgets


type BudgetsBudgetActionSubscriber struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#address BudgetsBudgetAction#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#subscription_type BudgetsBudgetAction#subscription_type}.
	SubscriptionType *string `field:"required" json:"subscriptionType" yaml:"subscriptionType"`
}

