// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawslblistener

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLbListenerMutualAuthenticationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbListenerMutualAuthenticationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbListenerMutualAuthenticationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbListenerMutualAuthenticationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbListenerMutualAuthenticationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLbListenerMutualAuthenticationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

