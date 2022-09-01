//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineGlobalFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmPatchBaselineGlobalFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

