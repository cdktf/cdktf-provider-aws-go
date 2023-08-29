// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dbproxy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbProxy) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validatePutAuthParameters(value interface{}) error {
	return nil
}

func (d *jsiiProxy_DbProxy) validatePutTimeoutsParameters(value *DbProxyTimeouts) error {
	return nil
}

func validateDbProxy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDbProxy_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDbProxy_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetDebugLoggingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetEngineFamilyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetIdleClientTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetRequireTlsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetRoleArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetVpcSecurityGroupIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_DbProxy) validateSetVpcSubnetIdsParameters(val *[]*string) error {
	return nil
}

func validateNewDbProxyParameters(scope constructs.Construct, id *string, config *DbProxyConfig) error {
	return nil
}

