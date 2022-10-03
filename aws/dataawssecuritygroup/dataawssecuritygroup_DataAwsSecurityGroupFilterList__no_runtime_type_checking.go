//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataawssecuritygroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsSecurityGroupFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsSecurityGroupFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsSecurityGroupFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsSecurityGroupFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsSecurityGroupFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsSecurityGroupFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsSecurityGroupFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

