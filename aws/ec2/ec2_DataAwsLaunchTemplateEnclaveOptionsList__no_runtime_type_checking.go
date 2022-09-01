//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLaunchTemplateEnclaveOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLaunchTemplateEnclaveOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateEnclaveOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateEnclaveOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateEnclaveOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLaunchTemplateEnclaveOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

