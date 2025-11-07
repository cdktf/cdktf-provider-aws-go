// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lbtargetgroup

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbTargetGroupTargetHealthStateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LbTargetGroupTargetHealthStateList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbTargetGroupTargetHealthStateList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbTargetGroupTargetHealthStateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbTargetGroupTargetHealthStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbTargetGroupTargetHealthStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbTargetGroupTargetHealthStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbTargetGroupTargetHealthStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

