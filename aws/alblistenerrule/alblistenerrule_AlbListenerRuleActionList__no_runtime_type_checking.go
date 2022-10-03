//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package alblistenerrule

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerRuleActionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerRuleActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerRuleActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

