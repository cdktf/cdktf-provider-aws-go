//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53RecordLatencyRoutingPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53RecordLatencyRoutingPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

