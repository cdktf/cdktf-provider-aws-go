// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vpcendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcEndpointSubnetConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointSubnetConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointSubnetConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpcEndpointSubnetConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VpcEndpointSubnetConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpcEndpointSubnetConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpcEndpointSubnetConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpcEndpointSubnetConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

