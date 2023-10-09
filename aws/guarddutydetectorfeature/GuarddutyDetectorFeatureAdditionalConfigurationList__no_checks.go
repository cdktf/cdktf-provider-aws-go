// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutydetectorfeature

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeatureAdditionalConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyDetectorFeatureAdditionalConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

