//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package budgets

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BudgetsBudgetCostFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BudgetsBudgetCostFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetCostFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetCostFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetCostFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetCostFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBudgetsBudgetCostFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

