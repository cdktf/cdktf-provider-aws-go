//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package alblistener

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionForwardTargetGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerDefaultActionForwardTargetGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

