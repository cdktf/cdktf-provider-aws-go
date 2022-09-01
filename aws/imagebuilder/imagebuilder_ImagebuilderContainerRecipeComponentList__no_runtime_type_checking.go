//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderContainerRecipeComponentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImagebuilderContainerRecipeComponentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

