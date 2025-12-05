// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ekscapability

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksCapabilityConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityConfigurationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksCapabilityConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

