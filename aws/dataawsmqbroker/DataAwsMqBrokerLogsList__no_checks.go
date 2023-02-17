//go:build no_runtime_type_checking

package dataawsmqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsMqBrokerLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerLogsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsMqBrokerLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

