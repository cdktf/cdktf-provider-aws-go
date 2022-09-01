//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package efs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEfsAccessPointPosixUserList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEfsAccessPointPosixUserList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEfsAccessPointPosixUserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEfsAccessPointPosixUserList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEfsAccessPointPosixUserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEfsAccessPointPosixUserListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

