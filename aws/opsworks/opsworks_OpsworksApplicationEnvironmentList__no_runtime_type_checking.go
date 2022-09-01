//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package opsworks

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpsworksApplicationEnvironmentList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OpsworksApplicationEnvironmentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OpsworksApplicationEnvironmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OpsworksApplicationEnvironmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OpsworksApplicationEnvironmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OpsworksApplicationEnvironmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOpsworksApplicationEnvironmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

