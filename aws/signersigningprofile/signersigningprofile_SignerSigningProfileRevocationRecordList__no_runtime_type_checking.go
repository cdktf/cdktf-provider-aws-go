//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package signersigningprofile

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignerSigningProfileRevocationRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignerSigningProfileRevocationRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignerSigningProfileRevocationRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

