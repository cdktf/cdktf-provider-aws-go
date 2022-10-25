//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketCorsRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketCorsRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketCorsRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketCorsRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketCorsRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketCorsRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketCorsRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

