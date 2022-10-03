//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cognitoresourceserver

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitoResourceServerScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CognitoResourceServerScopeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CognitoResourceServerScopeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CognitoResourceServerScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CognitoResourceServerScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CognitoResourceServerScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCognitoResourceServerScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

