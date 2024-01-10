// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lightsaildistribution

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDistributionLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LightsailDistributionLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDistributionLocationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDistributionLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

