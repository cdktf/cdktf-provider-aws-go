// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEcsServiceLoadBalancerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

