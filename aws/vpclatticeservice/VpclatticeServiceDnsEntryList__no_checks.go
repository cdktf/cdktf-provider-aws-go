//go:build no_runtime_type_checking

package vpclatticeservice

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpclatticeServiceDnsEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpclatticeServiceDnsEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpclatticeServiceDnsEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpclatticeServiceDnsEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpclatticeServiceDnsEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpclatticeServiceDnsEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

