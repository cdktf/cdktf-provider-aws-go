// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appstreamstack

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamStackAccessEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

