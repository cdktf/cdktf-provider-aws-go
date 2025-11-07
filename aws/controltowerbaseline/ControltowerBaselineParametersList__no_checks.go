// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package controltowerbaseline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ControltowerBaselineParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ControltowerBaselineParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ControltowerBaselineParametersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ControltowerBaselineParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ControltowerBaselineParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ControltowerBaselineParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ControltowerBaselineParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewControltowerBaselineParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

