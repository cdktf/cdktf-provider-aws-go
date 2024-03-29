// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package m2environment

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_M2EnvironmentStorageConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_M2EnvironmentStorageConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_M2EnvironmentStorageConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentStorageConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewM2EnvironmentStorageConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

