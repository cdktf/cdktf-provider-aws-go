//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsImagebuilderContainerRecipeComponentParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsImagebuilderContainerRecipeComponentParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderContainerRecipeComponentParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderContainerRecipeComponentParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderContainerRecipeComponentParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsImagebuilderContainerRecipeComponentParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

