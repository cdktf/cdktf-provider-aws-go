// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appflowflow

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppflowFlowTaskList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppflowFlowTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppflowFlowTaskList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppflowFlowTaskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppflowFlowTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppflowFlowTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppflowFlowTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppflowFlowTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

