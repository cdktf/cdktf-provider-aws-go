// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsalblistener

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbListenerDefaultActionForwardTargetGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbListenerDefaultActionForwardTargetGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionForwardTargetGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionForwardTargetGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionForwardTargetGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbListenerDefaultActionForwardTargetGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

