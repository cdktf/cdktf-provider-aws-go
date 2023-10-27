// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package albtargetgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbTargetGroupTargetHealthStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbTargetGroupTargetHealthStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

