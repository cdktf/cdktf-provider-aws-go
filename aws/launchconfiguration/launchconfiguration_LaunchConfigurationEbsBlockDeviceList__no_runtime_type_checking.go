//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package launchconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchConfigurationEbsBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchConfigurationEbsBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

