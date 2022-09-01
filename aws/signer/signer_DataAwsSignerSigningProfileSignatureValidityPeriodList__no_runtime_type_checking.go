//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package signer

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriodList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriodList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriodList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriodList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriodList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsSignerSigningProfileSignatureValidityPeriodListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

