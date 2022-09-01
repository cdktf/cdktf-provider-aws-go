//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package mq

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsMqBrokerConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsMqBrokerConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

