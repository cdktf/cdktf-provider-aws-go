//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package lblistener

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListenerDefaultActionList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbListenerDefaultActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbListenerDefaultActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbListenerDefaultActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

