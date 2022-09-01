//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricStreamStatisticsConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchMetricStreamStatisticsConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

