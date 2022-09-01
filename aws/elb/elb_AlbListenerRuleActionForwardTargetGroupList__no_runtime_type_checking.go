//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionForwardTargetGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerRuleActionForwardTargetGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

