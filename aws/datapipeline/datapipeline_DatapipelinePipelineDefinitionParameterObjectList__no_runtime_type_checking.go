//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datapipeline

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterObjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatapipelinePipelineDefinitionParameterObjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

