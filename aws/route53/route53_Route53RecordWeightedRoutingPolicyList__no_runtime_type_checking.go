//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53RecordWeightedRoutingPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53RecordWeightedRoutingPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

