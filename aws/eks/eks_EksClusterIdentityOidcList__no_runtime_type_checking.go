//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksClusterIdentityOidcList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksClusterIdentityOidcList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityOidcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityOidcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksClusterIdentityOidcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksClusterIdentityOidcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

