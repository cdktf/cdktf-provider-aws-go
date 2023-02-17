//go:build no_runtime_type_checking

package appconfigextension

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigExtensionActionPointList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionActionPointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppconfigExtensionActionPointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

