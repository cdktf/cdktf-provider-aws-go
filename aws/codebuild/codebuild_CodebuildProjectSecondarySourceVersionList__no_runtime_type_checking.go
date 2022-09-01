//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package codebuild

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourceVersionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodebuildProjectSecondarySourceVersionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

