//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmiCopyEbsBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmiCopyEbsBlockDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmiCopyEbsBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmiCopyEbsBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmiCopyEbsBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmiCopyEbsBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmiCopyEbsBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

