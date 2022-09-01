//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsNetworkmanagerLinkBandwidthList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsNetworkmanagerLinkBandwidthList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkmanagerLinkBandwidthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkmanagerLinkBandwidthList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkmanagerLinkBandwidthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsNetworkmanagerLinkBandwidthListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

