//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package budgets

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BudgetsBudgetActionSubscriberList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BudgetsBudgetActionSubscriberList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetActionSubscriberList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetActionSubscriberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetActionSubscriberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetActionSubscriberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBudgetsBudgetActionSubscriberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

