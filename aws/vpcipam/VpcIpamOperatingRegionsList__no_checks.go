//go:build no_runtime_type_checking

package vpcipam

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcIpamOperatingRegionsList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpcIpamOperatingRegionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpcIpamOperatingRegionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VpcIpamOperatingRegionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpcIpamOperatingRegionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpcIpamOperatingRegionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpcIpamOperatingRegionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

