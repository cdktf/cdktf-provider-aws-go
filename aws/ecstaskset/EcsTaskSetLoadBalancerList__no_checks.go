// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecstaskset

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsTaskSetLoadBalancerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsTaskSetLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsTaskSetLoadBalancerList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetLoadBalancerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetLoadBalancerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsTaskSetLoadBalancerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

