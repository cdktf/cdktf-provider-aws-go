//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleTimestreamList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleTimestreamList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleTimestreamList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleTimestreamList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleTimestreamList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleTimestreamList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleTimestreamListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

