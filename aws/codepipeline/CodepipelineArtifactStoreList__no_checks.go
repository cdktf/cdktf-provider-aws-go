//go:build no_runtime_type_checking

package codepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodepipelineArtifactStoreList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodepipelineArtifactStoreList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodepipelineArtifactStoreListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

