//go:build no_runtime_type_checking

package dataawsvpc

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsVpcFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsVpcFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsVpcFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

