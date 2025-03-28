// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package codepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodepipelineTriggerAllList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodepipelineTriggerAllList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodepipelineTriggerAllList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodepipelineTriggerAllList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodepipelineTriggerAllList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodepipelineTriggerAllList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodepipelineTriggerAllListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

