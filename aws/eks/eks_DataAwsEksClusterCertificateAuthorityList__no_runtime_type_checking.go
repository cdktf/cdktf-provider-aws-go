//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEksClusterCertificateAuthorityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEksClusterCertificateAuthorityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEksClusterCertificateAuthorityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEksClusterCertificateAuthorityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEksClusterCertificateAuthorityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEksClusterCertificateAuthorityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

