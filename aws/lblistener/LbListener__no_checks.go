// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lblistener

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListener) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_LbListener) validatePutDefaultActionParameters(value interface{}) error {
	return nil
}

func (l *jsiiProxy_LbListener) validatePutMutualAuthenticationParameters(value *LbListenerMutualAuthentication) error {
	return nil
}

func (l *jsiiProxy_LbListener) validatePutTimeoutsParameters(value *LbListenerTimeouts) error {
	return nil
}

func validateLbListener_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLbListener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLbListener_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLbListener_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetAlpnPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetCertificateArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetLoadBalancerArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetPortParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetProtocolParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetSslPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_LbListener) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func validateNewLbListenerParameters(scope constructs.Construct, id *string, config *LbListenerConfig) error {
	return nil
}

