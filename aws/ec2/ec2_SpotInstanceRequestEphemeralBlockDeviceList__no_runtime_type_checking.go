//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpotInstanceRequestEphemeralBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpotInstanceRequestEphemeralBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

