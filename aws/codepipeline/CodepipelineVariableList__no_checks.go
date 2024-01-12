// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package codepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodepipelineVariableList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodepipelineVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodepipelineVariableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodepipelineVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodepipelineVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodepipelineVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodepipelineVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodepipelineVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

