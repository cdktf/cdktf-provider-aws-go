//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2CoipPoolFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEc2CoipPoolFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

