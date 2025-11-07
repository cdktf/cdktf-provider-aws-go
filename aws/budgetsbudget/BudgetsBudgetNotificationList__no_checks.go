// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package budgetsbudget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BudgetsBudgetNotificationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BudgetsBudgetNotificationList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BudgetsBudgetNotificationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetNotificationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetNotificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetNotificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BudgetsBudgetNotificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBudgetsBudgetNotificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

