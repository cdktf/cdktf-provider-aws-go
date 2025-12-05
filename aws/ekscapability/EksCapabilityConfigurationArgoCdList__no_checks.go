// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ekscapability

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksCapabilityConfigurationArgoCdListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

