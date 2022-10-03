//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package mwaaenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MwaaEnvironmentLastUpdatedList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MwaaEnvironmentLastUpdatedList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMwaaEnvironmentLastUpdatedListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

