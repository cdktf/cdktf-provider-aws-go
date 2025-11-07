// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3directorybucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3DirectoryBucketLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3DirectoryBucketLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3DirectoryBucketLocationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3DirectoryBucketLocationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3DirectoryBucketLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3DirectoryBucketLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3DirectoryBucketLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3DirectoryBucketLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

