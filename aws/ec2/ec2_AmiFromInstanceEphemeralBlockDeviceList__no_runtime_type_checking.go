//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmiFromInstanceEphemeralBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmiFromInstanceEphemeralBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

