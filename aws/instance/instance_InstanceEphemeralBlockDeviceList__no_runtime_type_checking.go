//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package instance

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceEphemeralBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InstanceEphemeralBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InstanceEphemeralBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstanceEphemeralBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstanceEphemeralBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InstanceEphemeralBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInstanceEphemeralBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

