// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package signersigningjob

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignerSigningJobSignedObjectList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignerSigningJobSignedObjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobSignedObjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobSignedObjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignerSigningJobSignedObjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignerSigningJobSignedObjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

