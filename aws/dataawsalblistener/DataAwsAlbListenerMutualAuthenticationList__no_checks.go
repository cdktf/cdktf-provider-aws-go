// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsalblistener

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbListenerMutualAuthenticationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbListenerMutualAuthenticationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerMutualAuthenticationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerMutualAuthenticationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbListenerMutualAuthenticationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbListenerMutualAuthenticationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

