//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ekscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksClusterIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksClusterIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksClusterIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

