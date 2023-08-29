// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appstreamimagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamImageBuilderAccessEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamImageBuilderAccessEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

