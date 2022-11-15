package budgetsbudget


type BudgetsBudgetAutoAdjustData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget#auto_adjust_type BudgetsBudget#auto_adjust_type}.
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// historical_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget#historical_options BudgetsBudget#historical_options}
	HistoricalOptions *BudgetsBudgetAutoAdjustDataHistoricalOptions `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

