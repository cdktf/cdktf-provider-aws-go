//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53record

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53RecordFailoverRoutingPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53RecordFailoverRoutingPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

