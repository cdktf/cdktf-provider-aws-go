//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleElasticsearchList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleElasticsearchList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleElasticsearchList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleElasticsearchList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleElasticsearchList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleElasticsearchList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleElasticsearchListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

