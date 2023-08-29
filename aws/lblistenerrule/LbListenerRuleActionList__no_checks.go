// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lblistenerrule

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListenerRuleActionList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbListenerRuleActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbListenerRuleActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

