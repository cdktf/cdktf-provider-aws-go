// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package routetable

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteTable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validatePutRouteParameters(value interface{}) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validatePutTimeoutsParameters(value *RouteTableTimeouts) error {
	return nil
}

func validateRouteTable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRouteTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRouteTable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRouteTable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetPropagatingVgwsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetTagsAllParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_RouteTable) validateSetVpcIdParameters(val *string) error {
	return nil
}

func validateNewRouteTableParameters(scope constructs.Construct, id *string, config *RouteTableConfig) error {
	return nil
}

