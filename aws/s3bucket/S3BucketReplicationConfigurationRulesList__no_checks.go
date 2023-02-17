//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketReplicationConfigurationRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

