//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleKinesisList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleKinesisList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleKinesisList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleKinesisList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleKinesisList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleKinesisList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleKinesisListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

