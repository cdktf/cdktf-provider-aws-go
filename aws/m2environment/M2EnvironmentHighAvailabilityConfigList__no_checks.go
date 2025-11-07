// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package m2environment

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_M2EnvironmentHighAvailabilityConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewM2EnvironmentHighAvailabilityConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

