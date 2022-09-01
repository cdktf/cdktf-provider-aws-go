//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDbProxyAuthList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDbProxyAuthList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDbProxyAuthListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

