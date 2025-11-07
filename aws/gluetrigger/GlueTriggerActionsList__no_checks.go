// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gluetrigger

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueTriggerActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueTriggerActionsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueTriggerActionsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerActionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerActionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerActionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerActionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueTriggerActionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

