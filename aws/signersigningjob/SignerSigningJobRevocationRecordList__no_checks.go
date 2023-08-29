// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package signersigningjob

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignerSigningJobRevocationRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignerSigningJobRevocationRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobRevocationRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobRevocationRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobRevocationRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignerSigningJobRevocationRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

