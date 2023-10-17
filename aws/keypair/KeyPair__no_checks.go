// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keypair

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyPair) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateImportFromParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateKeyPair_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateKeyPair_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyPair_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateKeyPair_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetKeyNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetKeyNamePrefixParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetPublicKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_KeyPair) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func validateNewKeyPairParameters(scope constructs.Construct, id *string, config *KeyPairConfig) error {
	return nil
}

