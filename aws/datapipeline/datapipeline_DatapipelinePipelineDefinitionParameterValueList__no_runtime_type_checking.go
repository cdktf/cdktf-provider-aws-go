//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datapipeline

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatapipelinePipelineDefinitionParameterValueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatapipelinePipelineDefinitionParameterValueListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

