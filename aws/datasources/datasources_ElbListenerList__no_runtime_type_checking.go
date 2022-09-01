//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datasources

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElbListenerList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElbListenerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElbListenerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElbListenerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElbListenerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElbListenerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElbListenerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

