// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alblistenerrule

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleConditionQueryStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerRuleConditionQueryStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

