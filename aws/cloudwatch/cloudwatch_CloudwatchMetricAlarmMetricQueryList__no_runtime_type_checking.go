//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchMetricAlarmMetricQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchMetricAlarmMetricQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

