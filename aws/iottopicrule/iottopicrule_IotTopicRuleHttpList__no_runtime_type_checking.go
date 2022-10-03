//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iottopicrule

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleHttpList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleHttpList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleHttpList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleHttpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleHttpList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleHttpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleHttpListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

