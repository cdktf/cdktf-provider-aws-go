//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package imagebuilder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsImagebuilderImageOutputResourcesAmisList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsImagebuilderImageOutputResourcesAmisList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderImageOutputResourcesAmisList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderImageOutputResourcesAmisList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsImagebuilderImageOutputResourcesAmisList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsImagebuilderImageOutputResourcesAmisListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

