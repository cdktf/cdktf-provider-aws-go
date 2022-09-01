//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotTopicRuleRepublishList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotTopicRuleRepublishList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleRepublishList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleRepublishList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleRepublishList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotTopicRuleRepublishList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotTopicRuleRepublishListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

