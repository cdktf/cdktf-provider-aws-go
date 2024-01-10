// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package daxparametergroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaxParameterGroupParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DaxParameterGroupParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DaxParameterGroupParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DaxParameterGroupParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DaxParameterGroupParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DaxParameterGroupParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DaxParameterGroupParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDaxParameterGroupParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

