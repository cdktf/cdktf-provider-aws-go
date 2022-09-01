//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package amplify

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppCustomRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppCustomRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppCustomRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

