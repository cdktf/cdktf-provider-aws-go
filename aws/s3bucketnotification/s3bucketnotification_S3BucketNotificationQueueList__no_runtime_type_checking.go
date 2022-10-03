//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3bucketnotification

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketNotificationQueueList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketNotificationQueueList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationQueueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationQueueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationQueueList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketNotificationQueueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketNotificationQueueListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

