package budgets


type BudgetsBudgetActionDefinitionIamActionDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#policy_arn BudgetsBudgetAction#policy_arn}.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#groups BudgetsBudgetAction#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#roles BudgetsBudgetAction#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action#users BudgetsBudgetAction#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}
