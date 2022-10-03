//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package amplifyapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppProductionBranchList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppProductionBranchList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppProductionBranchList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppProductionBranchList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppProductionBranchList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppProductionBranchListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

