//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmPatchBaselineSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmPatchBaselineSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmPatchBaselineSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

