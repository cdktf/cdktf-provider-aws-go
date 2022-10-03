//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iottopicrule

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleSqsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleSqsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleSqsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleSqsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleSqsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleSqsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleSqsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

