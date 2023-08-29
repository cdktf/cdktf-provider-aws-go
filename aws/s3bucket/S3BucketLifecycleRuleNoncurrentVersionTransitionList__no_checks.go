// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionTransitionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketLifecycleRuleNoncurrentVersionTransitionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

