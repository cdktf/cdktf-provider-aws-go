//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleCloudwatchLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleCloudwatchLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

