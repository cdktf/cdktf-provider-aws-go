// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsalblistener

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbListenerDefaultActionRedirectList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbListenerDefaultActionRedirectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionRedirectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionRedirectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerDefaultActionRedirectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbListenerDefaultActionRedirectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

