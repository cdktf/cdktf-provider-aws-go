//go:build no_runtime_type_checking

package loadbalancerpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPolicyPolicyAttributeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerPolicyPolicyAttributeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

