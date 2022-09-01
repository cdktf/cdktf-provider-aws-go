//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package signer

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsSignerSigningJobRevocationRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

