//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elb

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListenerRuleConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbListenerRuleConditionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbListenerRuleConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbListenerRuleConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

