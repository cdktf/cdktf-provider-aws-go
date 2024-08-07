// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3objectcopy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ObjectCopyGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3ObjectCopyGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3ObjectCopyGrantList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3ObjectCopyGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3ObjectCopyGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3ObjectCopyGrantList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3ObjectCopyGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3ObjectCopyGrantListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

