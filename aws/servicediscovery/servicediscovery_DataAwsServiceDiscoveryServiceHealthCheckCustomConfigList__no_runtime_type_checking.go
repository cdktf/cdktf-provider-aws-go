//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package servicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsServiceDiscoveryServiceHealthCheckCustomConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsServiceDiscoveryServiceHealthCheckCustomConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsServiceDiscoveryServiceHealthCheckCustomConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsServiceDiscoveryServiceHealthCheckCustomConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsServiceDiscoveryServiceHealthCheckCustomConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsServiceDiscoveryServiceHealthCheckCustomConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

