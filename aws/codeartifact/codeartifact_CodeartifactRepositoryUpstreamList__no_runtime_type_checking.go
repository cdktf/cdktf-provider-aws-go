//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package codeartifact

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeartifactRepositoryUpstreamList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodeartifactRepositoryUpstreamList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryUpstreamList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryUpstreamList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryUpstreamList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryUpstreamList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodeartifactRepositoryUpstreamListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

