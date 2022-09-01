//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateBlockDeviceMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchTemplateBlockDeviceMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

