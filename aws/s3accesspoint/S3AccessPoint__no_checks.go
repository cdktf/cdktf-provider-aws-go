// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3accesspoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessPoint) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validatePutPublicAccessBlockConfigurationParameters(value *S3AccessPointPublicAccessBlockConfiguration) error {
	return nil
}

func (s *jsiiProxy_S3AccessPoint) validatePutVpcConfigurationParameters(value *S3AccessPointVpcConfiguration) error {
	return nil
}

func validateS3AccessPoint_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateS3AccessPoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3AccessPoint_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateS3AccessPoint_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetBucketParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetBucketAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessPoint) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewS3AccessPointParameters(scope constructs.Construct, id *string, config *S3AccessPointConfig) error {
	return nil
}

