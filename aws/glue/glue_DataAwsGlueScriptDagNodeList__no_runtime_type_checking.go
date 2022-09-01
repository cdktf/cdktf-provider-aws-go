//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsGlueScriptDagNodeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsGlueScriptDagNodeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsGlueScriptDagNodeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

