// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lambdafunction

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunction) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutCapacityProviderConfigParameters(value *LambdaFunctionCapacityProviderConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutDeadLetterConfigParameters(value *LambdaFunctionDeadLetterConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutDurableConfigParameters(value *LambdaFunctionDurableConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutEnvironmentParameters(value *LambdaFunctionEnvironment) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutEphemeralStorageParameters(value *LambdaFunctionEphemeralStorage) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutFileSystemConfigParameters(value *LambdaFunctionFileSystemConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutImageConfigParameters(value *LambdaFunctionImageConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutLoggingConfigParameters(value *LambdaFunctionLoggingConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutSnapStartParameters(value *LambdaFunctionSnapStart) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutTenancyConfigParameters(value *LambdaFunctionTenancyConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutTimeoutsParameters(value *LambdaFunctionTimeouts) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutTracingConfigParameters(value *LambdaFunctionTracingConfig) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validatePutVpcConfigParameters(value *LambdaFunctionVpcConfig) error {
	return nil
}

func validateLambdaFunction_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLambdaFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLambdaFunction_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLambdaFunction_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetArchitecturesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetCodeSigningConfigArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetFilenameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetFunctionNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetHandlerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetImageUriParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetKmsKeyArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetLayersParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetMemorySizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetPackageTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetPublishParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetPublishToParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetReplacementSecurityGroupIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetReplaceSecurityGroupsOnDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetReservedConcurrentExecutionsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetRoleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetRuntimeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetS3BucketParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetS3KeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetS3ObjectVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetSkipDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetSourceCodeHashParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetSourceKmsKeyArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunction) validateSetTimeoutParameters(val *float64) error {
	return nil
}

func validateNewLambdaFunctionParameters(scope constructs.Construct, id *string, config *LambdaFunctionConfig) error {
	return nil
}

