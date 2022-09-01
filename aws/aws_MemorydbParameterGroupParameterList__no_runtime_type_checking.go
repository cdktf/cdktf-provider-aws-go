//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbParameterGroupParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbParameterGroupParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbParameterGroupParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MemorydbParameterGroupParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbParameterGroupParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbParameterGroupParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbParameterGroupParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

