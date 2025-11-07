// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package evidentlyfeature

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureEvaluationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyFeatureEvaluationRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

