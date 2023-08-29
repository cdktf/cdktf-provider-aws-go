// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dbinstance

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbInstanceListenerEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbInstanceListenerEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbInstanceListenerEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbInstanceListenerEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbInstanceListenerEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbInstanceListenerEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

