// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServiceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceServiceConnectConfigurationServiceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

