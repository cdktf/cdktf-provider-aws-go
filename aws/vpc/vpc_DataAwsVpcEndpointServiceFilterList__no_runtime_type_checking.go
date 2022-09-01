//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcEndpointServiceFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsVpcEndpointServiceFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

