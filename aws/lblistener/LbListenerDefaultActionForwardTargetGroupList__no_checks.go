// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lblistener

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionForwardTargetGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbListenerDefaultActionForwardTargetGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

