//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AwsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAwsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetAssumeRoleParameters(val *AwsProviderAssumeRole) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetAssumeRoleWithWebIdentityParameters(val *AwsProviderAssumeRoleWithWebIdentity) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetDefaultTagsParameters(val *AwsProviderDefaultTags) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetEndpointsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetIgnoreTagsParameters(val *AwsProviderIgnoreTags) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetS3ForcePathStyleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetS3UsePathStyleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetSkipCredentialsValidationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetSkipGetEc2PlatformsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetSkipRegionValidationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetSkipRequestingAccountIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetUseDualstackEndpointParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsProvider) validateSetUseFipsEndpointParameters(val interface{}) error {
	return nil
}

func validateNewAwsProviderParameters(scope constructs.Construct, id *string, config *AwsProviderConfig) error {
	return nil
}

