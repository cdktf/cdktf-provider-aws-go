//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ami

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmiEphemeralBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmiEphemeralBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmiEphemeralBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

