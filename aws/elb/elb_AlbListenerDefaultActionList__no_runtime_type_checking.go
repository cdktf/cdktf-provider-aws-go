//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerDefaultActionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerDefaultActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerDefaultActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

