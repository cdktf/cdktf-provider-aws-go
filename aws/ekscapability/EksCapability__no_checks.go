// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ekscapability

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksCapability) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validatePutConfigurationParameters(value interface{}) error {
	return nil
}

func (e *jsiiProxy_EksCapability) validatePutTimeoutsParameters(value *EksCapabilityTimeouts) error {
	return nil
}

func validateEksCapability_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateEksCapability_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEksCapability_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateEksCapability_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetCapabilityNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetClusterNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetDeletePropagationPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetRoleArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_EksCapability) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewEksCapabilityParameters(scope constructs.Construct, id *string, config *EksCapabilityConfig) error {
	return nil
}

