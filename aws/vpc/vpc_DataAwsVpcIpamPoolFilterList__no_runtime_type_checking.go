//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsVpcIpamPoolFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsVpcIpamPoolFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

