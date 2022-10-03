//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceLoadBalancerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceLoadBalancerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceLoadBalancerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceLoadBalancerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

