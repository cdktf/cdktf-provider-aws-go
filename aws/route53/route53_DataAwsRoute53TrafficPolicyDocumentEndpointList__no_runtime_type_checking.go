//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsRoute53TrafficPolicyDocumentEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

