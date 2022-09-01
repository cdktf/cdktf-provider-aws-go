//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package backup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsBackupFrameworkControlInputParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsBackupFrameworkControlInputParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlInputParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlInputParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlInputParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsBackupFrameworkControlInputParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

