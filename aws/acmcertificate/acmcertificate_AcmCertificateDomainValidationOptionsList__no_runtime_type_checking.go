//go:build no_runtime_type_checking

package acmcertificate

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AcmCertificateDomainValidationOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AcmCertificateDomainValidationOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAcmCertificateDomainValidationOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

