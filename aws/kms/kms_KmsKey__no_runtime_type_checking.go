//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kms

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsKey) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KmsKey) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateKmsKey_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetBypassPolicyLockoutSafetyCheckParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetCustomerMasterKeySpecParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetDeletionWindowInDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetEnableKeyRotationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetIsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetKeyUsageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetMultiRegionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_KmsKey) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func validateNewKmsKeyParameters(scope constructs.Construct, id *string, config *KmsKeyConfig) error {
	return nil
}
