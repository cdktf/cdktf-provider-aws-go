//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketNotificationTopicList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketNotificationTopicList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationTopicList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationTopicList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationTopicList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationTopicList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketNotificationTopicListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

