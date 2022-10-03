//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glueuserdefinedfunction

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueUserDefinedFunctionResourceUrisList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueUserDefinedFunctionResourceUrisListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

