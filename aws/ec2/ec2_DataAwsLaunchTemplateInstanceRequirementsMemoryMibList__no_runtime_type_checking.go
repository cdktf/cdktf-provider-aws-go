//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsMemoryMibList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsMemoryMibList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsMemoryMibList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsMemoryMibList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsMemoryMibList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLaunchTemplateInstanceRequirementsMemoryMibListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

