//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MedialiveInputInputDevicesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MedialiveInputInputDevicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputInputDevicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputInputDevicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputInputDevicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputInputDevicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMedialiveInputInputDevicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

